package response

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//JsonResponse 通用返回
type JsonResponse struct {
	Code    int         `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

// PageData 分页数据
type PageData struct {
	PageNum int `json:"pageNum"`
	Total   int `json:"total"`

	Items interface{} `json:"items"`
}

//Json 返回 JSON 数据并退出当前 HTTP 执行函数。
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	if message == "" {
		message = GetMsg(code)
	} else {
		message = fmt.Sprintf("%v: %v", GetMsg(code), message)
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
