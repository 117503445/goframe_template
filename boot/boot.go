package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	_ "goframe_learn/packed"
)

func init() {
	LogBindEs()

	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// gres.Dump()

	InitDatabase()
}
