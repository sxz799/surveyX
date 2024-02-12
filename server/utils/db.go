package utils

import (
	"github.com/glebarez/sqlite"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	sqlType := config.SqlType
	dsn := config.SqlUrl
	switch sqlType {
	case "postgres":
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("postgres数据库连接失败。", err)
		}
	case "mysql":
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("mysql数据库连接失败。", err)
		}
	case "sqlite":
		var err error
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("sqlite数据库连接失败。", err)
		}
	}
	initDBTables()
}

func initDBTables() {
	_ = DB.AutoMigrate(entity.Survey{})
	_ = DB.AutoMigrate(entity.Question{})
	_ = DB.AutoMigrate(entity.Option{})
	_ = DB.AutoMigrate(entity.Answer{})
	_ = DB.AutoMigrate(entity.User{})
	initUser()
}

func initUser() {
	var u entity.User
	u.Nickname = "管理员"
	u.Username = "admin"
	u.Password = "admin"
	_ = DB.Create(&u)
}
