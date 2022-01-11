package main

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

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
