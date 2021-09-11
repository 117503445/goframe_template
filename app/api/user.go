package api

import (
	"goframe_template/app/dao"
	"goframe_template/app/model"
	"goframe_template/app/service"
	"goframe_template/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 用户API管理对象
var User = new(userApi)

type userApi struct{}

// @summary 用户注册
// @tags    user
// @accept json
// @produce json
// @Param b body model.UserServiceSignUpReq true "UserServiceSignUpReq"
// @router  /api/user [POST]
// @success 200 {object} response.JsonResponse
func (*userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.Json(r, response.Fail, "", err)
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.Json(r, response.Fail, "", err)
	}
	if err := service.User.SignUp(serviceReq); err != nil {
		response.Json(r, response.ErrorCreateFail, "", err)
	} else {
		response.Json(r, response.Success, "", nil)
	}
}

// @summary 用户获取自己的信息
// @tags    user
// @accept json
// @produce json
// @router  /api/user [GET]
// @Security JWT
// @success 200 {object} response.JsonResponse
func (*userApi) GetInfo(r *ghttp.Request) {
	user := new(model.User)
	if err := r.GetCtxVar("user").Struct(user); err != nil {
		response.Json(r, response.Fail, "", err)
	} else {
		roles := dao.GetRolesByUser(user)

		roleNames := g.Array{}
		for _, r := range roles {
			// g.Log().Line().Debug(r.Name)
			roleNames = append(roleNames, r.Name)
		}
		// g.Log().Line().Debug(dao.HasRole(user,"admin"))
		response.Json(r, response.Success, "", g.Map{"id": user.Id, "username": user.Username, "roles": roleNames})
	}
}
