package api

import (
    "DuckyGo/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

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
