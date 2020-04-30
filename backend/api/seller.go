package api

import (
    "DuckyGo/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

// SellerAddBook 卖家上传二手书信息接口
func SellerAddBook(c *gin.Context)  {
    var (
        userId  service.UserHeader
        reqBody service.SubSellerAddBookService
    )
    file, err := c.FormFile("cover")
    if err != nil {
        c.JSON(http.StatusOK, ErrorResponse(err))
        return
    } else {
        reqBody.Cover = service.UploadFile(*file)
    }
    file, err = c.FormFile("descp")
    if err != nil {
        c.JSON(http.StatusOK, ErrorResponse(err))
        return
    } else {
        reqBody.Descp = service.UploadFile(*file)
    }
    if err = c.ShouldBindHeader(&userId); err == nil {
        if err = c.ShouldBind(&reqBody); err == nil {
            serv := service.SellerAddBookService{Header: userId, Body:   reqBody}
            c.JSON(http.StatusOK, serv.AddBook())
        } else {
            c.JSON(http.StatusOK, ErrorResponse(err))
        }
    } else {
        c.JSON(http.StatusOK, ErrorResponse(err))
    }
}

// SellerShowBook 卖家查看自己售卖的二手书
func SellerShowBook(c *gin.Context)  {
    var header service.SellerShowBookService
    if err := c.ShouldBindHeader(&header); err == nil {
        c.JSON(http.StatusOK, header.ShowBook())
    } else {
        c.JSON(http.StatusOK, ErrorResponse(err))
    }
}

// SellerUpdateBook 卖家修改二手书信息
func SellerUpdateBook(c *gin.Context)  {
    var (
        userId  service.UserHeader
        reqBody service.SubSellerChangeBookService
    )
    file, err := c.FormFile("cover")
    if file == nil {
        // 如果file == nil，表示用户没有上传封面图片
        reqBody.Cover = nil
    } else if err != nil {
        c.JSON(http.StatusOK, ErrorResponse(err))
        return
    } else {
        reqBody.Cover = (*service.UploadFile)(file)
    }
    file, err = c.FormFile("descp")
    if file == nil {
        reqBody.Descp = nil
    } else if err != nil {
        c.JSON(http.StatusOK, ErrorResponse(err))
        return
    } else {
        reqBody.Descp = (*service.UploadFile)(file)
    }
    if err = c.ShouldBindHeader(&userId); err == nil {
        if err = c.ShouldBind(&reqBody); err == nil {
            serv := service.SellerChangeBookService{Header: userId, Body:   reqBody}
            c.JSON(http.StatusOK, serv.ChangeBook())
        } else {
            c.JSON(http.StatusOK, ErrorResponse(err))
        }
    } else {
        c.JSON(http.StatusOK, ErrorResponse(err))
    }
}

// SellerDelBook 卖家删除书籍信息
func SellerDelBook(ctx *gin.Context)  {
    var bookId  service.SubDelBookServiceBody
    if err := ctx.ShouldBind(&bookId); err == nil {
        serv := service.SellerDelBookService{Body: bookId}
        ctx.JSON(http.StatusOK, serv.DelBook())
    } else {
        ctx.JSON(http.StatusOK, ErrorResponse(err))
    }
}

// SellerShowOrder 卖家查看订单信息
func SellerShowOrder(ctx *gin.Context) {
    var serv service.SellerShowOrderService
    if err := ctx.ShouldBindHeader(&serv); err == nil {
        ctx.JSON(http.StatusOK, serv.ShowOrder())
    } else {
        ctx.JSON(http.StatusOK, ErrorResponse(err))
    }
}

// SellerPostCommodity 卖家发货
func SellerPostCommodity(ctx *gin.Context)  {
    var (
        header 	service.UserHeader
        body 	service.BodyPostCommodityService
    )
    if err := ctx.ShouldBindHeader(&header); err == nil {
        if err = ctx.ShouldBind(&body); err == nil {
            serv := service.PostCommodityService{Header: header, Body: body}
            ctx.JSON(http.StatusOK, serv.PostCommodity())
        } else {
            ctx.JSON(http.StatusOK, ErrorResponse(err))
        }
    } else {
        ctx.JSON(http.StatusOK, ErrorResponse(err))
    }
}