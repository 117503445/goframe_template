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
	// g.Log().Line().Debug(link)

	linkWithoutDbName := strings.Join(strings.Split(strings.Split(link, "/")[0], ":")[:][1:], ":") + "/"
	dbName := strings.Split(link, "/")[1]

	// g.Log().Line().Debug(linkWithoutDbName)

	sqlDB, err := sql.Open("mysql", linkWithoutDbName)
	if err != nil {
		g.Log().Line().Panic(err)
	}
	defer sqlDB.Close()

	query := fmt.Sprintf("SELECT * FROM information_schema.SCHEMATA where SCHEMA_NAME='%v'", dbName)
	// g.Log().Debug(query)

	result, err := sqlDB.Query(query)
	if err != nil {
		g.Log().Line().Panic(err)
	}

	isDbExists := false
	for result.Next() {
		isDbExists = true // 只要查到了一条记录,就说明数据库存在
		break
	}
	// g.Log().Debug(isDbExists)
	if !isDbExists {
		g.Log().Info(fmt.Sprintf("CREATE DATABASE %v", dbName))
		_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName))
		if err != nil {
			g.Log().Line().Panic(err)
		}
	}

}

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})

	InitDatabase()
}
