package api

import (
	"github.com/gogf/gf/net/ghttp"
	"goframe_learn/library/response"
)

// @summary 返回 Hello
// @tags    hello
// @produce json
// @router  / [GET]
// @success 200 {object} response.JsonResponse
func Hello(r *ghttp.Request) {
	response.Json(r, 0, "","hello")
}
