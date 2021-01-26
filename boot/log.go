package boot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"goframe_learn/library/elasticsearch"
	"strings"
)

type EsLogWriter struct {
	logger *glog.Logger
}

func (w *EsLogWriter) Write(p []byte) (n int, err error) {
	body, _ := json.Marshal(g.Map{"body": string(p)}) // todo parse log

	req := esapi.IndexRequest{
		Index:   g.Cfg().GetString("elasticsearch.index"),
		Body:    strings.NewReader(string(body)),
		Refresh: "true",
	}

	// Perform the request with the client.
	if _, err := req.Do(context.Background(), elasticsearch.Es); err != nil {
		fmt.Println(err)
	}

	return w.logger.Write(p)
}

func LogBindEs() {
	if g.Cfg().GetBool("elasticsearch.enabled") {
		g.Log().Line().Debug("LogBindEs")
		g.Log().SetWriter(&EsLogWriter{logger: glog.DefaultLogger()})
		g.Log().Line().Debug("test log")
	}

}
