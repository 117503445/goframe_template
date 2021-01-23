package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"goframe_learn/app/model"
	"goframe_learn/app/service"
	"goframe_learn/library/response"
)

// 用户API管理对象
var User = new(userApi)

type userApi struct{}

func (a *userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.Json(r, response.Fail, "", err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.Json(r, response.Fail, "", err.Error())
	}
	if err := service.User.SignUp(serviceReq); err != nil {
		response.Json(r, response.ErrorCreateFail, "", err.Error())
	} else {
		response.Json(r, 0, "", nil)
	}
}
