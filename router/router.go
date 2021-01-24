package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe_learn/app/api"
	"goframe_learn/app/middleware"
)

func middlewareAuth(r *ghttp.Request) {
	middleware.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Hello)
		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.Group("/user", func(group *ghttp.RouterGroup) {
				group.POST("/login", middleware.Auth.LoginHandler)
				group.POST("/", api.User.SignUp)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middlewareAuth)
					group.GET("/", api.User.GetInfo)
				})
			})
		})
	})
}
