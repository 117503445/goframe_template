package main

import (
	_ "goframe_template/boot"
	_ "goframe_template/router"

	"github.com/gogf/gf/frame/g"
)

// @title       goframe_template API
// @version     1.6.2
// @description `goframe_template` 117503445 的 goframe 学习/模板项目 api

// @contact.name 117503445
// @contact.url http://www.117503445.top
// @contact.email t117503445@gmail.com

// @license.name GNU GPL 3.0

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

func main() {
	s := g.Server()
	s.SetDumpRouterMap(g.Cfg().GetBool("server.DumpRouterMap"))
	s.Run()
}
