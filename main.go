package main

import (
	_ "goframe_learn/boot"
	_ "goframe_learn/router"

	"github.com/gogf/gf/frame/g"
)

// @title       goframe_learn API
// @version     1.1.3
// @description `goframe_learn` 117503445 的 goframe 学习/模板项目 api
// @schemes     http

// @contact.name 117503445
// @contact.url http://www.117503445.top
// @contact.email t117503445@gmail.com

// @license.name GNU GPL 3.0

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

func main() {
	g.Server().Run()
}
