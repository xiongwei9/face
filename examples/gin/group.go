// a demo for generated router-group
package main

import (
	"github.com/gin-gonic/gin"
)

// func ping(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "ok in router-group",
// 	})
// }

type UserInfo struct {
	OpenId   string
	UserName string
}

func NewMethod(fn func(r *gin.Context, target string) (*UserInfo, error)) func(r *gin.Context) {
	return func(r *gin.Context) {
		target := r.Param("target")
		fn(r, target)
	}
}

func getUserInfo(r *gin.Context, target string) (*UserInfo, error) {
	return nil, nil
}

func BaseRouteGroup(r *gin.RouterGroup) {
	r.GET("/getUserInfo", NewMethod(getUserInfo))
}
