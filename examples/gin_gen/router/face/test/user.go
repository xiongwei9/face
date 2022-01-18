package test

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

type Sex int32

const (
	SexMale   Sex = 1
	SexFemale Sex = 2
)

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Age      int32  `json:"age"`
}

type UserInfo struct {
	OpenId   string            `json:"openId"`
	UserName string            `json:"userName"`
	Sex      Sex               `json:"sex"`
	Favorite []string          `json:"favorite"`
	Fri      map[string]string `json:"fri"`
}

type TestService struct{}

func (s *TestService) getUserInfo(c *gin.Context, params *string) (*UserInfo, error) {
	resp := &UserInfo{
		OpenId:   "xiongwei9",
		UserName: "xiongwei",
		Sex:      SexMale,
		Favorite: []string{"computer"},
		Fri:      nil,
	}
	return resp, nil
}
func (s *TestService) login(c *gin.Context, params *LoginRequest) (*string, error) {
	var resp *string = nil // TODO
	return resp, nil
}

func TestService_RouteGroup(r *gin.RouterGroup) {
	s := &TestService{}

	r.GET("/api/getUserInfo", NewMethod(s.getUserInfo))
	r.POST("/api/login", NewMethod(s.login))

}
