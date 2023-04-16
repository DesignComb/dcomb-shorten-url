package server

import (
	"github.com/gin-gonic/gin"
	"main/pkg/jwt"
	"main/pkg/res"
)

// GetUserInfo GetUserInfo
func GetUserInfo(c *gin.Context) {
	token, err := c.Cookie(jwt.Key)
	if err != nil {
		res.Success(c, gin.H{
			"is_login":       false,
			"user_id":        "",
			"google_user_id": "",
			"user_email":     "",
			"user_name":      "",
			"user_picture":   "",
		})
		return
	}

	userId, googleUserId, email, name, picture, err := jwt.ParseToken(token)
	if err != nil {
		res.Success(c, gin.H{
			"is_login":       false,
			"user_id":        "",
			"google_user_id": "",
			"user_email":     "",
			"user_name":      "",
			"user_picture":   "",
		})
		return
	}

	res.Success(c, gin.H{
		"is_login":       true,
		"user_id":        userId,
		"google_user_id": googleUserId,
		"user_email":     email,
		"user_name":      name,
		"user_picture":   picture,
	})
}
