package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"goframe_learn/app/dao"
	"goframe_learn/app/model"
	"goframe_learn/library/response"
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
