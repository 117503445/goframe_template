package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe_learn/app/api"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Hello)
		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.ALL("/user", api.User)
		})
	})
}
