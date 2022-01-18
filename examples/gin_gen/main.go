package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiongwei9/face/examples/gin_gen/router/face/test"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	baseGroup := r.Group("/")
	test.TestService_RouteGroup(baseGroup)

	r.Run(":8888")
}
