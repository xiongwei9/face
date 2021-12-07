// a demo for generated router-group
package main

import (
	"errors"
	"net/http"
	"reflect"

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

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}

func NewMethod(fn interface{}) func(r *gin.Context) {
	funcType := reflect.TypeOf(fn)
	funcValue := reflect.ValueOf(fn)
	if funcType.Kind() != reflect.Func {
		panic("The route handler must be a function")
	}
	paramRef := funcType.In(1)
	paramType := paramRef.Elem()

	return func(r *gin.Context) {
		param := reflect.New(paramType).Interface()
		if err := r.ShouldBind(param); err != nil {
			r.String(http.StatusBadRequest, err.Error())
			return
		}

		resp := funcValue.Call(valOf(r, param))
		if len(resp) != 2 {
			r.String(http.StatusInternalServerError, "system error")
			return
		}
		if !resp[1].IsNil() {
			msg := resp[1].MethodByName("Error").Call(nil)
			r.String(http.StatusInternalServerError, msg[0].String())
			return
		}

		result := resp[0].Interface()
		r.JSON(http.StatusOK, result)
	}
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
