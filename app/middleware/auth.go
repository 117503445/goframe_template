package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"goframe_template/app/model"
	"goframe_template/app/service"
)

func NeedRole(role string) func(*ghttp.Request) {
	return func(r *ghttp.Request) {
		user := new(model.User)
		if err := r.GetCtxVar("user").Struct(user); err != nil {
			panic(err)
		} else {
			if service.UserRole.HasRole(r.Context(), user, role) {
				r.Middleware.Next()
			} else {
				panic(err)
			}
		}
	}
}
