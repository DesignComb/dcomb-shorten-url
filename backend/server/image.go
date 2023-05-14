package server

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"main/pkg/res"
	"main/utils"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func UploadImage(c *gin.Context) {
	// 取得上傳的file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 打開上傳的file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	fileName := utils.RandomURL(8) + filepath.Ext(file.Filename)

	url := "https://sgcv33ycze.execute-api.ap-northeast-1.amazonaws.com/v1/short-url-img-bucket/upload/" + fileName
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, url, src)
	if err != nil {
		res.SystemError(c, gin.H{"error": err.Error()}, "upload failed")
		return
	}
	// 設定header
	req.Header.Set("Content-Type", file.Header.Get("Content-Type"))
	resp, err := client.Do(req)
	if err != nil {
		res.SystemError(c, gin.H{"error": err.Error()}, "upload failed")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.SystemError(c, gin.H{"error": err.Error()}, "upload failed")
		return
	}
	if len(body) != 0 {
		res.SystemError(c, nil, "upload failed")
		return
	}
	defer resp.Body.Close()

	res.Success(c, nil)
}
