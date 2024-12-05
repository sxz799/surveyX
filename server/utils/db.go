package utils

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) (db *gorm.DB) {
	sqlType := config.SqlType
	dsn := config.SqlUrl
	switch sqlType {
	case "postgres":
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("postgres数据库连接失败。", err)
		}
	case "mysql":
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("mysql数据库连接失败。", err)
		}
	case "sqlite":
		var err error
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("sqlite数据库连接失败。", err)
		}
	}
	initDBTables(db)
	return db
}

func initDBTables(db *gorm.DB) {
	_ = db.AutoMigrate(entity.Survey{})
	_ = db.AutoMigrate(entity.Question{})
	_ = db.AutoMigrate(entity.Option{})
	_ = db.AutoMigrate(entity.Answer{})
	_ = db.AutoMigrate(entity.User{})
	initUser(db)
}

func initUser(db *gorm.DB) {

	users := []entity.User{
		{Nickname: "管理员", Username: "admin", Password: "admin"},
		{Nickname: "33", Username: "33", Password: "33"},
	}

	_ = db.Create(&users)
}
