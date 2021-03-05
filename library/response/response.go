package response

import (
	"fmt"
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
		data = d.Error()
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
