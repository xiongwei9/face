package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok in router-group",
	})
}

func BaseRouteGroup(r *gin.RouterGroup) {
	r.GET("/ping", ping)
}
