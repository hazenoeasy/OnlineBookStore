package service

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type SubSellerChangeBookService struct {
	// 下面不是必须要填写的字段，所以全部设置为指针
	// 当相应字段为nil时，表示不需要修改该字段
	Id     int     `form:"id" json:"id" binding:"required"`
	Title  *string `form:"title" json:"title"`
	Author *string `form:"author" json:"author"`
	Price  *int    `form:"price" json:"price"`
	Num    *int    `form:"num" json:"num"`
	Kind   *string `form:"kind" json:"kind"`
	Cover  *UploadFile
	Descp  *UploadFile
}

type SellerChangeBookService struct {
	Header UserHeader
	Body   SubSellerChangeBookService
}

// 注意：此版本中，当更新封面和描述图片时，没有删除原文件，这可能会影响性能
func (p *SellerChangeBookService) ChangeBook() serializer.Response {
	var (
		path                     = os.Getenv("STATIC_PATH")
		userId                   = strconv.Itoa(p.Header.UserId)
		now                      = strconv.FormatInt(time.Now().Unix(), 10)
		coverName, descpName     string
		coverNotify, descpNotify chan error
	)
	// 保存封面
	if p.Body.Cover != nil {
		coverName = filepath.Join(path, userId+"cover"+now+p.Body.Cover.Filename)
		coverNotify = make(chan error, 1)
		go p.Body.Cover.SaveUploadFile(coverName, coverNotify)
	}
	// 保存描述图片
	if p.Body.Descp != nil {
		descpName = filepath.Join(path, userId+"descp"+now+p.Body.Descp.Filename)
		descpNotify = make(chan error, 1)
		go p.Body.Descp.SaveUploadFile(descpName, descpNotify)
	}
	// 修改数据库记录
	modify := make(map[string]interface{})

	if p.Body.Title != nil {
		modify["title"] = *(p.Body.Title)
	}
	if p.Body.Author != nil {
		modify["author"] = *(p.Body.Author)
	}
	if p.Body.Price != nil {
		modify["price"] = *(p.Body.Price)
	}
	if p.Body.Num != nil {
		modify["num"] = *(p.Body.Num)
	}
	if p.Body.Kind != nil {
		modify["kind"] = *(p.Body.Kind)
	}
	if p.Body.Cover != nil {
		modify["cover_url"] = filepath.ToSlash(coverName)
	}
	if p.Body.Descp != nil {
		modify["descp_url"] = filepath.ToSlash(descpName)
	}

	tx := model.DB.Begin()
	if err := tx.Model(&model.Book{BookId: p.Body.Id}).Updates(modify).Error; err != nil {
		tx.Rollback()
		// 数据库更新失败，如果文件上传成功，则将文件删除掉
		if coverNotify != nil && <-coverNotify == nil {
			os.Remove(coverName)
		}
		if descpNotify != nil && <-descpNotify == nil {
			os.Remove(descpName)
		}
		return serializer.Response{
			Code: serializer.DBWriteErr,
			Data: err.Error(),
			Msg:  "数据库：图书信息更新失败",
		}
	}
	// 如果数据库更新成功，而文件上传失败了
	if coverNotify != nil {
		if err := <-coverNotify; err != nil {
			tx.Rollback()
			os.Remove(descpName)
			return serializer.Response{
				Code: serializer.FileSaveErr,
				Data: err,
				Msg:  "封面保存失败",
			}
		}
	}
	if descpNotify != nil {
		if err := <-descpNotify; err != nil {
			tx.Rollback()
			os.Remove(coverName)
			return serializer.Response{
				Code: serializer.FileSaveErr,
				Data: err,
				Msg:  "描述图片保存失败",
			}
		}
	}
	tx.Commit()
	return serializer.Response{
		Code: serializer.OpSuccess,
		Data: nil,
		Msg:  "ok",
	}
}
