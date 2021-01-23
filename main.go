package main

import (
	_ "goframe_learn/boot"
	_ "goframe_learn/router"

	"github.com/gogf/gf/frame/g"
)

// @title       goframe_learn API
// @version     1.0.6
// @description `goframe_learn` 117503445 的 goframe 学习/模板项目 api
// @schemes     http
func main() {
	g.Server().Run()
}
