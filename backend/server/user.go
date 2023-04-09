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
			"is_login":     false,
			"user_id":      "",
			"user_email":   "",
			"user_name":    "",
			"user_picture": "",
		})
		return
	}

	id, email, name, picture, err := jwt.ParseToken(token)
	if err != nil {
		res.Success(c, gin.H{
			"is_login":     false,
			"user_id":      "",
			"user_email":   "",
			"user_name":    "",
			"user_picture": "",
		})
		return
	}

	res.Success(c, gin.H{
		"is_login":     true,
		"user_id":      id,
		"user_email":   email,
		"user_name":    name,
		"user_picture": picture,
	})
}
