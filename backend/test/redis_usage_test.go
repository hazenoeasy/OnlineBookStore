package test

import (
    "fmt"
    "github.com/go-redis/redis"
    "strconv"
    "strings"
    "testing"
)

func connect() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    _, err := client.Ping().Result()
    if err != nil {
        panic(fmt.Errorf("ping error[%s]\n", err.Error()))
    }
    return client
}

func TestZset(t *testing.T)  {
    fmt.Printf("开始测试ZSet\n")

    client := connect()
    defer client.Close()

    // 开始测试zset
    zsetKey := "hot_rank"
    ranking := []redis.Z{
        redis.Z{Score: 100.0, Member: "钟南山"},
        redis.Z{Score: 80.0, Member: "林医生"},
        redis.Z{Score: 70.0, Member: "王医生"},
        redis.Z{Score: 75.0, Member: "张医生"},
        redis.Z{Score: 59.0, Member: "叶医生"},
    }
    client.ZAdd(zsetKey, ranking...)
    list, _ := client.ZRevRangeWithScores(zsetKey, 0, 2).Result()
    fmt.Printf("热度前三名：%v\n", list)
}

func TestHashSet(t *testing.T)  {
    fmt.Printf("开始测试HashSet\n")

    client := connect()
    defer client.Close()

    // 开始测试HashSET
    data := map[string]interface{}{
        "book_id":  1,
        "title":    "高数",
        "author":   "章星明",
        "kind":     "大学数学",
        "price":    53,
        "salesnum": 12,
        "num":      100,
        "cover":    "static/12355.png",
        "descp":    "static/af2323.png",
    }
    key := strings.Join([]string{"book_id",
        strconv.Itoa(data["book_id"].(int))},":")
    client.HMSet(key, data)

    result := client.HGetAll(key).Val()
    for k, v := range result {
        fmt.Printf("key: %s, value: %v\n", k, v)
    }
}