package utils

import (
	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"code": status,
		"data": data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"code": status,
		"data": gin.H{
			"error": message,
		},
	})
}
