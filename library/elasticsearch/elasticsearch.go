package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gogf/gf/frame/g"
)

var Es *elasticsearch.Client

func init() {
	if len(g.Cfg().GetString("elasticsearch.index")) > 0 {
		g.Log().Line().Info("elasticsearch init")
		cfg := elasticsearch.Config{
			Username: g.Cfg().GetString("elasticsearch.username"),
			Password: g.Cfg().GetString("elasticsearch.password"),
		}
		var err error
		if Es, err = elasticsearch.NewClient(cfg); err != nil {
			g.Log().Line().Panic(err)
		}
		g.Log().Line().Info("elasticsearch init finish")
	}

}
