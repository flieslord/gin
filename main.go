package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)
func main() {
	r := gin.New()
	fmt.Println("服务器开始运行在8080端口")
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":8080")
}
