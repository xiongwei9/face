# Face 接口生成

## 生成目录结构

````
/facegen
    /namespace/dirname
        /service1.go
        /service2.go
    /utils.go # NewMethod函数
````

## 生成代码示例

结构体：

````golang
type UserInfoResponse struct {
	OpenId   string `json:"openId"`
	UserName string `json:"useName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserInfoRequest struct {
	OpenId string `form:"openId"`
}
````

接口：

````golang
func getUserInfo(r *gin.Context, params *UserInfoRequest) (*UserInfoResponse, error) {
	userInfo := &UserInfoResponse{}
	return userInfo, nil
}

func ServiceName_RouteGroup(r *gin.RouterGroup) {
	r.GET("/getUserInfo", NewMethod(getUserInfo))
}
````
