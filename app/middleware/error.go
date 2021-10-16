package middleware

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe_template/library/response"
)

func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()

		var data string

		if g.Cfg().GetBool("server.returnErrStack") {
			data = gerror.Stack(err)
		} else {
			data = ""
		}

		if err := r.Response.WriteJson(response.JsonResponse{
			Code:    gerror.Code(err).Code(),
			Message: gerror.Code(err).Message(),
			Data:    data,
		}); err != nil {
			g.Log().Line().Error(err)
		}

	}
}
