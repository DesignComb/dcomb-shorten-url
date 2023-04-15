package middleware

import (
	"github.com/gin-gonic/gin"
	"main/pkg/jwt"
	"net/http"
)

// Auth Auth
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取得token
		token, err := c.Cookie(jwt.Key)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"result":     false,
				"error_code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		// 解析token 取得會員的資料
		id, email, name, picture, err := jwt.ParseToken(token)
		if err != nil || id == "" || email == "" || name == "" || picture == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"result":     false,
				"error_code": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		// 把值傳到下一層
		c.Set("user_id", id)
		c.Set("user_email", email)
		c.Set("user_name", name)
		c.Set("user_picture", picture)

		c.Next()
	}
}