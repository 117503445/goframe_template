package api

import (
	jwt "github.com/gogf/gf-jwt"
	"goframe_template/app/model"
	"goframe_template/app/service"
	"goframe_template/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// User 用户API管理对象
var User = new(userApi)

type userApi struct{}

// SignUp
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
	if err := service.User.SignUp(r.Context(), serviceReq); err != nil {
		response.Json(r, response.ErrorCreateFail, "", err)
	} else {
		response.Json(r, response.Success, "", nil)
	}
}

// Login is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// Check error (e) to determine the appropriate error message.
// @summary 用户登录
// @tags    user
// @accept json
// @produce json
// @Param b body model.UserApiLoginReq true "UserApiLoginReq"
// @router  /api/user/login [POST]
// @success 200 {object} response.JsonResponse
func Login(r *ghttp.Request) (interface{}, error) {
	var (
		apiReq     *model.UserApiLoginReq
		serviceReq *model.UserServiceLoginReq
	)
	if err := r.Parse(&apiReq); err != nil {
		return "", err
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		return "", err
	}

	if user := service.User.GetUserByUsernamePassword(r.Context(), serviceReq); user != nil {
		return user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

// GetInfo
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
		roles := service.UserRole.GetRolesByUser(r.Context(), user)

		roleNames := g.Array{}
		for _, r := range roles {
			roleNames = append(roleNames, r.Name)
		}
		response.Json(r, response.Success, "", g.Map{"id": user.Id, "username": user.Username, "roles": roleNames})
	}
}
