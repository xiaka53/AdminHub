package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/AdminHub/middleware"
	"github.com/xiaka53/AdminHub/redis"
	"mime/multipart"
)

func Upload(c *gin.Context) {
	var (
		headerFile *multipart.FileHeader
		info       upload_info
		ok         bool
		err        error
	)
	token := c.GetHeader("token")
	ok = (&redis.Upload{
		Uuid: token,
	}).GetUuid()
	if !ok {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	headerFile, err = c.FormFile("file")
	if err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	info.Name = headerFile.Filename
	err = c.SaveUploadedFile(headerFile, "files/"+info.Name)
	if err != nil {
		middleware.ResponseError(c, middleware.AddErr, errors.New(""))
		return
	}
	info.Url = "http://127.0.0.1:81/files/" + info.Name
	middleware.ResponseSuccess(c, info)
}

type upload_info struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
