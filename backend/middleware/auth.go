package middleware

import (
	"github.com/gin-gonic/gin"
	"main/pkg/jwt"
	"main/pkg/res"
)

// Auth Auth
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取得token
		token, err := c.Cookie(jwt.Key)
		if err != nil {
			res.Unauthorized(c, nil, "not logged in")
			c.Abort()
			return
		}
		// 解析token 取得會員的資料
		userId, googleUserId, email, name, picture, err := jwt.ParseToken(token)
		if err != nil || userId == "" || googleUserId == "" || email == "" || name == "" || picture == "" {
			res.Unauthorized(c, nil, "not logged in")
			c.Abort()
			return
		}

		// 把值傳到下一層
		c.Set("user_id", userId)
		c.Set("google_user_id", googleUserId)
		c.Set("user_email", email)
		c.Set("user_name", name)
		c.Set("user_picture", picture)

		c.Next()
	}
}
