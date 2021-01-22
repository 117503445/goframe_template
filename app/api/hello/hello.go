package hello

import (
	"github.com/gogf/gf/net/ghttp"
)

// @summary 返回 Hello
// @tags    Hello
// @produce plain
// @router  / [GET]
// @success 200 {string} string "Hello World!"
func Hello(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
