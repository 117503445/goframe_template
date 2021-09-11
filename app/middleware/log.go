package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"goframe_template/library/elasticsearch"
	"io/ioutil"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
	requestBody := ""
	if requestBodyBytes, err := ioutil.ReadAll(r.Request.Body); err == nil && len(requestBodyBytes) > 0 {
		requestBody = string(requestBodyBytes) + "\n"
	}
	responseBody := string(r.Response.Buffer())
	//g.Log().Info(requestMethod)
	//g.Log().Info(url)
	//g.Log().Info(requestBody)
	//g.Log().Info(statusCode)
	//g.Log().Info(responseBody)
	g.Log().Info(fmt.Sprintf("[HTTP] %v %v\n%v%v %v", requestMethod, url, requestBody, statusCode, responseBody))

	body, _ := json.Marshal(g.Map{"requestMethod": requestMethod, "url": url.Path, "requestBody": requestBody, "statusCode": statusCode, "responseBody": responseBody})

	if g.Cfg().GetBool("logger.elasticsearch") {
		req := esapi.IndexRequest{
			Index:   g.Cfg().GetString("elasticsearch.index"),
			Body:    bytes.NewReader(body),
			Refresh: "true",
		}

		if _, err := req.Do(context.Background(), elasticsearch.Es); err != nil {
			fmt.Println(err)
		}
	}

}
