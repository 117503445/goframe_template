package middleware

import (
	"goframe_template/app/model"
	"goframe_template/app/service"
	"goframe_template/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func NeedRole(role string) func(*ghttp.Request) {
	return func(r *ghttp.Request) {
		user := new(model.User)
		if err := r.GetCtxVar("user").Struct(user); err != nil {
			response.ErrorResp(r, err)
		} else {
			if service.UserRole.HasRole(r.Context(), user, role) {
				r.Middleware.Next()
			} else {
				response.ErrorResp(r, err)
			}
		}
	}
}
