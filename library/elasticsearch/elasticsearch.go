package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gogf/gf/frame/g"
)

var Es *elasticsearch.Client

func init() {
	if g.Cfg().Get("elasticsearch.enabled").(bool) {
		g.Log().Line().Info("elasticsearch init")
		cfg := elasticsearch.Config{
			// ...
			Username: g.Cfg().Get("elasticsearch.username").(string),
			Password: g.Cfg().Get("elasticsearch.password").(string),
		}
		var err error
		if Es, err = elasticsearch.NewClient(cfg); err != nil {
			g.Log().Line().Panic(err)
		}
		g.Log().Line().Info("elasticsearch init finish")
	}

}
