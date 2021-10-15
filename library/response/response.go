package response

import (
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//JsonResponse 返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

//Json 返回JSON数据并退出当前HTTP执行函数。
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	if message == "" {
		message = GetMsg(code)
	} else {
		message = fmt.Sprintf("%v: %v", GetMsg(code), message)
	}

	switch d := data.(type) {
	case error:
		if gerror.Code(d).Code() != -1 {
			// https://goframe.org/pages/viewpage.action?pageId=3671864#:~:text=%E7%A0%81%E4%BF%A1%E6%81%AF%E6%97%B6%EF%BC%8C-,%E8%AF%A5%E6%96%B9%E6%B3%95%E8%BF%94%E5%9B%9E-1,-%EF%BC%8C%E5%90%A6%E5%88%99%E8%BF%94%E5%9B%9Eerror
			// 受 goframe 管理的 error
			// 根据错误中的信息进行返回

			message = gerror.Code(d).Message()
			code = gerror.Code(d).Code()
		}

		if g.Cfg().GetBool("server.returnErrStack") {
			data = gerror.Stack(d)
		} else {
			g.Log().Line().Error(d)
			data = nil
		}
	}

	if err := r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}); err != nil {
		g.Log().Line().Error(err)
	}

	r.Exit()
}

// ErrorResp 返回错误
func ErrorResp(r *ghttp.Request, err error) {
	Json(r, Error, "", err)
}
