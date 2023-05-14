package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    http.StatusOK,
		//"message": message,
		"data":    payload,
	})
}

func BadRequest(c *gin.Context, payload interface{}, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":     "error",
		"code":       http.StatusBadRequest,
		"message":    message,
		"data":       payload,
	})
}

func SystemError(c *gin.Context, payload interface{}, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":     "error",
		"code":       http.StatusBadRequest,
		"message":    message,
		"data":       payload,
	})
}

func Unauthorized(c *gin.Context, payload interface{}, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":     "error",
		"code":       http.StatusUnauthorized,
		"message":    message,
		"data":       payload,
	})
}
