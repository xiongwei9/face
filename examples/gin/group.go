// a demo for generated router-group
package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

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

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}

func NewMethod(fn interface{}) func(r *gin.Context) {
	// Test
	{
		var getUserInfoFn interface{} = getUserInfo
		funcType := reflect.TypeOf(getUserInfoFn)
		funcValue := reflect.ValueOf(getUserInfoFn)
		if funcType.Kind() != reflect.Func {
			panic("The route handler must be a function")
		}

		paramRef := funcType.In(1)
		fmt.Printf("%v, %v", funcValue, paramRef)
	}

	funcType := reflect.TypeOf(fn)
	funcValue := reflect.ValueOf(fn)
	if funcType.Kind() != reflect.Func {
		panic("The route handler must be a function")
	}

	i := funcType.NumIn()
	fmt.Println(i)
	paramRef := funcType.In(1)
	paramType := paramRef.Elem()
	// param := reflect.New(paramType).Interface()
	return func(r *gin.Context) {
		param := reflect.New(paramType).Interface()
		if err := r.ShouldBind(&param); err != nil {
			r.String(http.StatusBadRequest, err.Error())
		}

		resp := funcValue.Call(valOf(r, param))
		if len(resp) != 2 {
			r.String(http.StatusInternalServerError, "system error")
		}
		if resp != nil {
			msg := resp[1].MethodByName("Error").Call(nil)
			r.String(http.StatusInternalServerError, msg[0].String())
			return
		}
		r.JSON(http.StatusOK, resp[0])
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
