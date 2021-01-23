package boot

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	_ "goframe_learn/packed"
	"strings"
)

//InitDatabase Create database if not exists
func InitDatabase() {
	link := g.Cfg().Get("database.link").(string)
	// g.Log().Line().Info(link)

	linkWithoutDbName := strings.Join(strings.Split(strings.Split(link, "/")[0], ":")[:][1:], ":") + "/"
	dbName := strings.Split(link, "/")[1]

	// g.Log().Line().Info(linkWithoutDbName)

	sqlDB, err := sql.Open("mysql", linkWithoutDbName)
	if err != nil {
		g.Log().Line().Panic(err)
	}
	defer sqlDB.Close()

	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName))
	if err != nil {
		g.Log().Line().Panic(err)
	}
}

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})

	InitDatabase()
}
