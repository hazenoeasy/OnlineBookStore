package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
)

type SubSellerChangeBookService struct {
    // 下面不是必须要填写的字段，所以全部设置为指针
    // 当相应字段为nil时，表示不需要修改该字段
    Id      int          `form:"id" json:"id" binding:"required"`
    Title   *string      `form:"title" json:"title"`
    Author  *string      `form:"author" json:"author"`
    Price   *int         `form:"price" json:"price"`
    Num     *int         `form:"num" json:"num"`
    Kind    *string      `form:"kind" json:"kind"`
    Cover   *UploadFile
    Descp   *UploadFile
}

type SellerChangeBookService struct {
    Header  UserHeader
    Body    SubSellerChangeBookService
}

func (p *SellerChangeBookService) ChangeBook() serializer.Response {
    // 保存封面
    var (
        path    = os.Getenv("STATIC_PATH")
        userId  = strconv.Itoa(p.Header.UserId)
        now     = strconv.FormatInt(time.Now().Unix(), 10)
        coverName, descpName        string
        coverNotify, descpNotify    chan string
    )
    if p.Body.Cover != nil {
        coverName   = filepath.Join(path, userId + "cover" + now + p.Body.Cover.Filename)
        coverNotify = make(chan string, 1)
        go p.Body.Cover.SaveUploadFile(coverName, coverNotify)
    }
    // 保存描述图片
    if p.Body.Descp != nil {
        descpName   = filepath.Join(path, userId + "descp" + now + p.Body.Descp.Filename)
        descpNotify = make(chan string, 1)
        go p.Body.Descp.SaveUploadFile(descpName, descpNotify)
    }
    // 修改数据库记录
    modify := make(map[string]interface{})
    switch {
    case p.Body.Title != nil:
        modify["title"] = p.Body.Title
    case p.Body.Author != nil:
        modify["author"] = p.Body.Author
    case p.Body.Price  != nil:
        modify["price"] = p.Body.Price
    case p.Body.Num != nil:
        modify["num"] = p.Body.Num
    case p.Body.Kind != nil:
        modify["kind"] = p.Body.Kind
    case p.Body.Cover != nil:
        modify["cover"] = strings.ReplaceAll(coverName, `\`, "/")
    case p.Body.Descp != nil:
        modify["descp"] = strings.ReplaceAll(descpName, `\`, "/")
    }
    book := model.Book{ BookId: p.Header.UserId }
    if err := model.DB.Model(&book).Updates(modify).Error; err != nil {
        // 如果更新了文件，并且文件成功保存，则将其删除
        switch {
        case coverNotify != nil && <-coverNotify == "":
            os.Remove(coverName)
        case descpNotify != nil && <-descpNotify == "":
            os.Remove(descpName)
        }
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err.Error(),
            Msg:  "数据库：图书信息更新失败",
        }
    }
    // TODO: 处理文件保存失败，但数据库更新成功后处理部分
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}