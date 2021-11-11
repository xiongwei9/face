// a demo for generated router-group
package main

import (
	"errors"
	"net/http"

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
	Email    string
	Phone    string
}

type UserInfoRequest struct {
	OpenId string `form:"openId" binding:"required"`
}

func NewMethod(fn func(r *gin.Context, params UserInfoRequest) (*UserInfo, error)) func(r *gin.Context) {
	return func(r *gin.Context) {
		var params UserInfoRequest
		if err := r.ShouldBind(&params); err != nil {
			r.String(http.StatusBadRequest, err.Error())
		}

		data, err := fn(r, params)
		if err != nil {
			r.String(http.StatusInternalServerError, err.Error())
			return
		}
		r.JSON(http.StatusOK, data)
	}
}

func getUserInfo(r *gin.Context, params UserInfoRequest) (*UserInfo, error) {
	if params.OpenId == "" {
		return nil, errors.New("params error")
	}
	userInfo := &UserInfo{
		UserName: params.OpenId,
		OpenId:   params.OpenId,
		Email:    params.OpenId + "@foxmail.com",
		Phone:    "-",
	}
	return userInfo, nil
}

func BaseRouteGroup(r *gin.RouterGroup) {
	r.GET("/getUserInfo", NewMethod(getUserInfo))
}
