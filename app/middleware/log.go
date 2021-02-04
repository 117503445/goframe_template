package middleware

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
)

func HttpLog(r *ghttp.Request) {
	r.Middleware.Next()
	//errStr := ""
	//if err := r.GetError(); err != nil {
	//	errStr = err.Error()
	//}
	url := r.URL
	statusCode := r.Response.Status
	requestMethod := r.Request.Method
	requestBody := "empty request body"
	if requestBodyBytes, err := ioutil.ReadAll(r.Request.Body); err == nil && len(requestBodyBytes) > 0 {
		requestBody = string(requestBodyBytes)
	}
	responseBody := string(r.Response.Buffer())
	//g.Log().Info(requestMethod)
	//g.Log().Info(url)
	//g.Log().Info(requestBody)
	//g.Log().Info(statusCode)
	//g.Log().Info(responseBody)
	g.Log().Info(fmt.Sprintf("[HTTP] %v %v\n%v\n%v %v", requestMethod, url, requestBody, statusCode, responseBody))
}
