package middleware

import (
	"goframe_template/app/dao"
	"goframe_template/app/model"
	"goframe_template/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func NeedRole(role string) func(*ghttp.Request) {
	return func(r *ghttp.Request) {
		user := new(model.User)
		if err := r.GetCtxVar("user").Struct(user); err != nil {
			response.Json(r, response.ErrorAuthCheckTokenFail, "", err)
		} else {
			if dao.HasRole(user, role) {
				r.Middleware.Next()
			} else {
				response.Json(r, response.ErrorAuthRoleFail, "", nil)
			}
		}
	}
}
