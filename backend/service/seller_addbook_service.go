package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// 处理文件上传的服务
type UploadFile multipart.FileHeader

// 保存上传的文件
// 输入文件的保存路径path（包括文件名），路径必须存在，否则出错
// 保存结果通过rtv返回，如果一切正常，则rtv返回nil，否则rtv返回错误信息
func (p *UploadFile) SaveUploadFile(path string, rtv chan<- error) {
	fh := (*multipart.FileHeader)(p)
	file, err := fh.Open()
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		rtv <- err
		return
	}
	upload := NewFileService(path, fh.Size, file)
	if upload.Upload() != nil {
		upload.UploadRollBack()
		rtv <- errors.New("文件上传失败")
	} else {
		upload.UploadCommit()
		rtv <- nil
	}
}

type SubSellerAddBookService struct {
	Title  string `form:"title" json:"title" binding:"required"`
	Author string `form:"author" json:"author" binding:"required"`
	Price  int    `form:"price" json:"price" binding:"required"`
	Num    int    `form:"num" json:"num" binding:"required"`
	Kind   string `form:"kind" json:"kind" binding:"required"`
	Cover  UploadFile
	Descp  UploadFile
}

type SellerAddBookService struct {
	Header UserHeader
	Body   SubSellerAddBookService
}

func (s *SellerAddBookService) AddBook() serializer.Response {
	// 先把图书封面和描述图书保存在应用服务器上
	var (
		path   = os.Getenv("STATIC_PATH")
		userId = strconv.Itoa(s.Header.UserId)
		now    = strconv.FormatInt(time.Now().Unix(), 10)
		coverName 	= filepath.Join(path, userId+"cover"+now+s.Body.Cover.Filename)
		descpName 	= filepath.Join(path, userId+"descp"+now+s.Body.Descp.Filename)
		descNotify 	= make(chan error, 1)
		coverNotify = make(chan error, 1)
	)
	// 保存封面和描述图片
	go s.Body.Cover.SaveUploadFile(coverName, coverNotify)
	go s.Body.Descp.SaveUploadFile(descpName, descNotify)

	// 图书信息写入数据库
	book := model.Book{
		SalesManId: s.Header.UserId,
		Title:      s.Body.Title,
		Author:     s.Body.Author,
		Price:      s.Body.Price,
		Num:        s.Body.Num,
		SalesNum:   0,
		Kind:       s.Body.Kind,
		CoverUrl:   filepath.ToSlash(coverName),
		DescpUrl:   filepath.ToSlash(descpName),
	}
	tx := model.DB.Begin()
	if err := tx.Create(&book).Error; err != nil {
		tx.Rollback()
		// 如果数据库保存失败，而文件成功保存，则将文件删除
		if <-coverNotify == nil {
			os.Remove(coverName)
		}
		if <-descNotify == nil {
			os.Remove(descpName)
		}
		return serializer.Response{
			Code: serializer.DBWriteErr,
			Data: err.Error(),
			Msg:  "数据库：保存图书信息失败",
		}
	}
	// 如果数据库保存成功，而文件保存失败，则回滚数据库
	if coverMsg := <-coverNotify; coverMsg != nil {
		tx.Rollback()
		os.Remove(descpName)
		return serializer.Response{
			Code: serializer.FileSaveErr,
			Data: coverMsg,
			Msg:  "封面保存失败",
		}
	}
	if descpMsg := <-descNotify; descpMsg != nil {
		tx.Rollback()
		os.Remove(coverName)
		return serializer.Response{
			Code: serializer.FileSaveErr,
			Data: descpMsg,
			Msg:  "描述图片保存失败",
		}
	}
	tx.Commit()
	return serializer.Response{
		Code: serializer.OpSuccess,
		Data: nil,
		Msg:  "ok",
	}
}
