package boot

import (
	"context"
	"database/sql"
	"fmt"
	"goframe_template/app/dao"
	"goframe_template/app/model"
	"goframe_template/library"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
)

//InitDatabase Create database if not exists
func InitDatabase() {
	link := g.Cfg().GetString("database.link")
	// g.Log().Line().Debug(link)

	linkWithoutDbName := strings.Join(strings.Split(strings.Split(link, "/")[0], ":")[:][1:], ":") + "/"
	dbName := strings.Split(link, "/")[1]

	// g.Log().Line().Debug(linkWithoutDbName)

	sqlDB, err := sql.Open("mysql", linkWithoutDbName)
	if err != nil {
		g.Log().Line().Panic(err)
	}
	defer sqlDB.Close()

	isForceCreate := g.Cfg().GetBool("database.forceCreate")
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
		g.Log().Line().Info(fmt.Sprintf("create database %v", dbName))

		if _, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName)); err != nil {
			g.Log().Line().Panic(err)
		}

		// g.Log().Line().Debug(createSQL)
		sqlMyDB, err := sql.Open("mysql", linkWithoutDbName+dbName+"?multiStatements=true")
		if err != nil {
			g.Log().Line().Panic(err)
		}

		arraySqlPath := g.Cfg().GetArray("database.sqlOnCreate")
		for _, path := range arraySqlPath {
			g.Log().Line().Debug("run ", path)
			sqlText := gfile.GetContents(path.(string))
			if _, err = sqlMyDB.Exec(sqlText); err != nil {
				g.Log().Line().Panic(err)
			}
		}

		randomAdminPasswordOnCreate := g.Cfg().GetBool("database.randomAdminPasswordOnCreate")
		if randomAdminPasswordOnCreate {
			adminPassword := library.RandStringRunes(12)
			if cipher, err := model.EncryptPassword(adminPassword); err != nil {
				g.Log().Line().Panic(err)
			} else {
				if err = gfile.PutContents("./tmp/password/admin.txt", adminPassword); err != nil {
					g.Log().Line().Error(err)
				}
				var user *model.User

				if err := dao.User.Ctx(context.TODO()).Where("username", "admin").Scan(&user); err != nil {
					g.Log().Line().Error(err)
				} else {
					user.Password = cipher
					if _, err = dao.User.Ctx(context.TODO()).Where("id", user.Id).Update(user); err != nil {
						g.Log().Line().Error(err)
					}
				}
			}
		}
	}
}
