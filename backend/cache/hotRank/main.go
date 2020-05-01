package hotRank

import (
    "DuckyGo/cache"
    "DuckyGo/model"
    "fmt"
    "github.com/go-redis/redis"
    "os"
    "strconv"
)

var HOT *HotRank

type HotRank struct {
    redis           *redis.Client
    rankKey         string
    objKeyPrefix    string
    limit           int64
}

func (this *HotRank) View(start, end int64) (books []model.Book) {
     // 保证排行榜和热榜书籍信息一致，防止查询为空
    this.redis.Watch(func(tx *redis.Tx) error {
        bookIdS := tx.ZRevRange(this.rankKey, start, end).Val()
        // 根据bookID去Redis Cache寻找书籍详细信息
        for i := 0; i < len(bookIdS); i++ {
            redisData := tx.HGetAll(this.getKeyFromBookIdStr(bookIdS[i])).Val()
            books = append(books, NewModelBookFormRedis(redisData))
        }
        return nil
    }, this.rankKey)
    return
}

func (this *HotRank) IsExisted(bookid int) bool {
    return this.redis.ZRank(this.rankKey, strconv.Itoa(bookid)).Err() == nil
}

func (this *HotRank) Add(book *model.Book) {
    // 保证排行榜和书籍信息同步更新，防止可能的查询为空
    this.redis.Watch(func(tx *redis.Tx) error {
        tx.ZAdd(this.rankKey, redis.Z{
            Score:  float64(book.SalesNum),
            Member: book.BookId,
        })
        data := NewRedisBookFromModel(book)
        tx.HMSet(this.getKeyFromBookId(book.BookId), data)
        return nil
    }, this.rankKey)

    // 如果热榜长度超出限制，则删除榜尾和它对应的Cache数据
    if this.redis.ZCard(this.rankKey).Val() > this.limit {
        this.RemoveLast()
    }
}

func (this *HotRank) getKeyFromBookId(bookid int) string {
    return this.objKeyPrefix + strconv.Itoa(bookid)
}
func (this *HotRank) getKeyFromBookIdStr(bookid string) string {
    return this.objKeyPrefix + bookid
}

func (this *HotRank) Update(bookid string, incr float64) {
    this.redis.ZIncrBy(this.rankKey, incr, bookid)
    // 更新销量和库存
    key := this.getKeyFromBookIdStr(bookid)
    this.redis.HIncrByFloat(key, "salesnum", incr)
    this.redis.HIncrByFloat(key, "num", -incr)
}

func (this *HotRank) RemoveLast()  {
    id := this.redis.ZRange(this.rankKey, 0, 0).Val()[0]
    this.redis.ZRem(this.rankKey, id)
    this.redis.Del(this.getKeyFromBookIdStr(id))
}

// 初始化HOT RANK服务
// 必须在redis和MySQL初始化后，调用此函数
func InitHotRank()  {
    HOT = &HotRank{
        redis:        cache.RedisClient,
        rankKey:      os.Getenv("HOT_RANK_NAME"),
        objKeyPrefix: "book_id:",
        limit:        100,
    }
    // 从数据库中取得TOP100的数据
    var books []model.Book
    if err := model.DB.Where("salesnum > 0").Order("salesnum DESC").
            Limit(HOT.limit).Find(&books).Error; err != nil {
            panic(fmt.Errorf("InitHotRank() err: %v", err))
    }
    // 将数据加入Hot Rank中
    for i := 0; i < len(books); i++ {
        HOT.Add(&books[i])
    }
}

type RedisBook          map[string]interface{}
type RedisBookString    map[string]string

func NewRedisBookFromModel(book *model.Book) RedisBook {
    return RedisBook{
        "book_id":  book.BookId,
        "title":    book.Title,
        "author":   book.Author,
        "kind":     book.Kind,
        "price":    book.Price,
        "salesnum": book.SalesNum,
        "num":      book.Num,
        "cover":    book.CoverUrl,
        "descp":    book.DescpUrl,
    }
}

func NewModelBookFormRedis(book RedisBookString) model.Book {
    var (
    	id, _       = strconv.Atoi(book["book_id"])
    	price, _    = strconv.Atoi(book["price"])
    	num, _      = strconv.Atoi(book["num"])
    	salesnum, _ = strconv.Atoi(book["salesnum"])
    )
    return model.Book{
        BookId:     id,
        Title:      book["title"],
        Author:     book["author"],
        Price:      price,
        Num:        num,
        SalesNum:   salesnum,
        Kind:       book["kind"],
        CoverUrl:   book["cover"],
        DescpUrl:   book["descp"],
    }
}