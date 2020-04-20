package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
    "mime/multipart"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
)

// 处理文件上传的服务
type UploadFile multipart.FileHeader

// 保存上传的文件
// 输入文件的保存路径path（包括文件名），路径必须存在，否则出错
// 保存结果通过rtv返回，如果一切正常，则rtv返回""，否则rtv返回错误信息
func (p *UploadFile) SaveUploadFile(path string, rtv chan<-string) {
    fh := (*multipart.FileHeader)(p)
    file, err := fh.Open()
    defer file.Close()
    if err != nil {
        rtv <- err.Error()
        return
    }
    // 将上传文件内容读出，之后写入本地文件中
    buf := make([]byte, fh.Size)
    if _, err = file.Read(buf); err != nil {
        rtv <- err.Error()
        return
    }
    out, err := os.Create(path)
    defer out.Close()
    if err != nil {
        rtv <- err.Error()
        return
    }
    if _, err = out.Write(buf); err != nil {
        rtv <- err.Error()
    } else {
        rtv <- ""
    }
}

type SubSellerAddBookService struct {
    Title   string      `form:"title" json:"title" binding:"required"`
    Author  string      `form:"author" json:"author" binding:"required"`
    Price   int         `form:"price" json:"price" binding:"required"`
    Num     int         `form:"num" json:"num" binding:"required"`
    Kind    string      `form:"kind" json:"kind" binding:"required"`
    Cover   UploadFile
    Descp   UploadFile
}

type SellerAddBookService struct {
    Header  UserHeader
    Body    SubSellerAddBookService
}

func (s *SellerAddBookService) AddBook() serializer.Response {
    // 先把图书封面和描述图书保存在应用服务器上
    path :=os.Getenv("STATIC_PATH")
    // 如果路径不存在，则创建路径
    _, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            os.MkdirAll(path, 0755)
        } else {
            return serializer.Response{
                Code: serializer.FileSaveErr,
                Data: err.Error(),
                Msg:  "系统错误（文件保存路径错误）",
            }
        }
    }
    // 保存封面
    var (
        userId  = strconv.Itoa(s.Header.UserId)
        now     = strconv.FormatInt(time.Now().Unix(), 10)
    )
    coverName   := filepath.Join(path, userId + "cover" + now + s.Body.Cover.Filename)
    coverNotify := make(chan string, 1)
    go s.Body.Cover.SaveUploadFile(coverName, coverNotify)
    // 保存描述图片
    descpName := filepath.Join(path, userId + "descp" + now + s.Body.Descp.Filename)
    descpNotity := make(chan string, 1)
    go s.Body.Descp.SaveUploadFile(descpName, descpNotity)

    // 图书信息写入数据库
    book := model.Book{
        SalesManId: s.Header.UserId,
        Title:      s.Body.Title,
        Author:     s.Body.Author,
        Price:      s.Body.Price,
        Num:        s.Body.Num,
        SalesNum:   0,
        Kind:       s.Body.Kind,
        CoverUrl:   strings.ReplaceAll(coverName, `\`, "/"),
        DescpUrl:   strings.ReplaceAll(descpName, `\`, "/"),
    }
    if err = model.DB.Create(&book).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err.Error(),
            Msg:  "数据库：保存图书信息失败",
        }
    }
    coverMsg := <- coverNotify
    descpMsg := <- descpNotity
    if coverMsg != "" || descpMsg != "" {
        return serializer.Response{
            Code: serializer.FileSaveErr,
            Data: "封面：" + coverMsg + "," + "描述图片" + descpMsg,
            Msg:  "图书封面/描述图片保存失败",
        }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}