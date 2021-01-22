package main

import (
	_ "goframe_learn/boot"
	_ "goframe_learn/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
