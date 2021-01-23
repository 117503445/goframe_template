package boot

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gres"
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

	isForceCreate := g.Cfg().Get("database.forceCreate").(bool)
	if isForceCreate {
		query := fmt.Sprintf("DROP DATABASE IF EXISTS %v", dbName)
		if _, err = sqlDB.Exec(query); err != nil {
			g.Log().Line().Panic(err)
		}
	}

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

		if _, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName)); err != nil {
			g.Log().Line().Panic(err)
		}

		createSQL := string(gres.Get("create.sql./create.sql").Content())
		// g.Log().Line().Debug(createSQL)
		sqlMyDB, err := sql.Open("mysql", linkWithoutDbName+dbName+"?multiStatements=true")
		if err != nil {
			g.Log().Line().Panic(err)
		}

		if _, err = sqlMyDB.Exec(createSQL); err != nil {
			g.Log().Line().Panic(err)
		}

		g.DB().Table("roles").Data(g.List{{"name": "admin"}, {"name": "user"}}).Save()

		g.DB().Table("users").Data(g.List{{"username": "admin", "password": "$2a$12$XLiKy7M77cY56.1aE9IxDeKONAHbz1Z0pE7IOmzRfpjTHKiMZYsjG"}}).Save()

		g.DB().Table("user_role").Data(g.List{{"user_id": "1", "role_id": "1"}, {"user_id": "1", "role_id": "2"}}).Save()

	}
}

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// gres.Dump()

	InitDatabase()
}
