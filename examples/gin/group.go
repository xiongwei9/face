// a demo for generated router-group
package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// func ping(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "ok in router-group",
// 	})
// }

type UserInfoResponse struct {
	OpenId   string `json:"openId"`
	UserName string `json:"useName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserInfoRequest struct {
	OpenId string `form:"openId"`
}

func getUserInfo(r *gin.Context, params *UserInfoRequest) (*UserInfoResponse, error) {
	if params.OpenId == "" {
		return nil, errors.New("params error")
	}
	userInfo := &UserInfoResponse{
		OpenId:   params.OpenId,
		UserName: params.OpenId,
		Email:    params.OpenId + "@foxmail.com",
		Phone:    "-",
	}
	return userInfo, nil
}

func BaseRouteGroup(r *gin.RouterGroup) {
	r.GET("/getUserInfo", NewMethod(getUserInfo))
}
