package server

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"main/model"
	"main/pkg/res"
	"main/utils"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
)

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

const s3UploadUrl = "https://sgcv33ycze.execute-api.ap-northeast-1.amazonaws.com/v1/short-url-img-bucket/upload/"
const s3ReadUrl = "https://short-url-img-bucket.s3.ap-northeast-1.amazonaws.com/upload/"

func UploadImage(c *gin.Context) {
	// 驗證user
	userId := c.Param("userId")
	var loginUserId uint64
	if len(userId) > 0 {
		value, exists := c.Get("user_id")
		if value != userId {
			res.Unauthorized(c, nil, "no permission")
			c.Abort()
			return
		}
		if exists {
			str, ok := value.(string)
			if ok {
				loginUserId, _ = strconv.ParseUint(str, 10, 64)
			}
		}
	}

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
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, s3UploadUrl+fileName, src)
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

	var image model.Image
	image.Uri = fileName
	image.UserId = loginUserId
	image, err = model.CreateImage(image)

	res.Success(c, gin.H{
		"id":     image.ID,
		"userId": image.UserId,
		"url":    s3ReadUrl + image.Uri,
	})
}
