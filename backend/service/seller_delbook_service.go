package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
    "errors"
    "os"
    "path/filepath"
)

type SubDelBookServiceBody struct {
    Id  int     `form:"id" json:"id" binding:"required"`
}

type SellerDelBookService struct {
    Header  UserHeader
    Body    SubDelBookServiceBody
}

func (this *SellerDelBookService) DelBook() serializer.Response {
    // 从数据库中找到书籍对应的其他文件资源位置，例如封面图片和描述图片
    book := model.Book{BookId: this.Body.Id}
    if err := model.DB.Select("book_id, cover_url, descp_url").Find(&book).Error;
        err != nil {
            return serializer.Response{
                Code: serializer.DBReadErr,
                Data: err,
                Msg:  "数据库：查询书籍失败",
            }
    }
    // 如果没有找到书籍信息
    if book.CoverUrl == "" || book.DescpUrl == "" {
        return serializer.Response{
            Code: serializer.DBReadErr,
            Data: errors.New("SellerDelBookService.DelBook: 未找到相关书籍"),
            Msg:  "书籍不存在",
        }
    }

    // 删除对应的文件资源 和 数据库记录（此处的数据库记录为必须删除项）
    // 查询和删除可以被打断，不需要事务化
    os.Remove(filepath.FromSlash(book.CoverUrl))
    os.Remove(filepath.FromSlash(book.DescpUrl))
    if err := model.DB.Delete(&book).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err,
            Msg:  "数据库：删除图书错误",
        }
    }

    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}
