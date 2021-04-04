package boot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gregex"
	"goframe_learn/library"
	"goframe_learn/library/elasticsearch"
)

type EsLogWriter struct {
	logger *glog.Logger
}

type EsLogBody struct {
	raw   string
	time  string
	level string
	line  string
	info  string
}

func makeEsLogBody(raw string) []byte {
	// fmt.Println(raw)

	time := raw[0:23]
	level := raw[25:29]

	p := library.StrIndex(raw, ":", 4) // 处理是否有行号的问题

	line := ""
	info := ""
	if p != -1 {
		line = raw[31:p]
		info = raw[p+2:]
	} else {
		info = raw[31:]
	}

	body, _ := json.Marshal(g.Map{"raw": raw, "time": time, "level": level, "line": line, "info": info})
	return body
}
func (w *EsLogWriter) Write(p []byte) (n int, err error) {

	if !gregex.IsMatch(`\[[A-Z]{4}\]`, p) {
		// 如果没有 [INFO] [DEBU]
		// 不发送至 ES
		return w.logger.Write(p)
	}

	body := makeEsLogBody(string(p))

	req := esapi.IndexRequest{
		Index:   g.Cfg().GetString("elasticsearch.index"),
		Body:    bytes.NewReader(body),
		Refresh: "true",
	}

	if _, err := req.Do(context.Background(), elasticsearch.Es); err != nil {
		fmt.Println(err)
	}

	return w.logger.Write(p)
}

func LogBindEs() {
	if g.Cfg().GetBool("logger.elasticsearch") {
		//g.Log().Line().Debug("LogBindEs")
		g.Log().SetWriter(&EsLogWriter{logger: glog.DefaultLogger()})
		//g.Log().Debug("test log")
		//g.Log().Line().Debug("test log")
	}

}
