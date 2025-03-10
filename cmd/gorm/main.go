package main

import (
	"fmt"

	"pastebin/biz/dal/model"
	"pastebin/biz/dal/mysql"
	"pastebin/conf"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	mysqldb "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var c = conf.GetConf()

func main() {
	// connect to mysql manually to check and create database
	dsn := "%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local&tls=true"
	db, err := gorm.Open(mysqldb.Open(fmt.Sprintf(dsn, c.MySQL.Username, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port)),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	var count int
	dbName := conf.GetConf().Hertz.Service
	db.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&count)
	if count == 0 {
		hlog.Warn("Database not found, creating database")
		db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", dbName))
	}

	// migrate the database
	mysql.Init()

	err = mysql.DB.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(model.AllModels...)
	if err != nil {
		panic(err)
	}
}
