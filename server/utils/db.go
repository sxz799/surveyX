package utils

import (
	"github.com/spf13/viper"
	"github.com/sxz799/surveyX/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	sqlType := viper.GetString("db.sqlType")
	switch sqlType {
	case "postgres":
		dsn := viper.GetString("db.postgresUrl")
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("postgres数据库连接失败。", err)
		}
	case "mysql":
		dsn := viper.GetString("db.mysqlUrl")
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("mysql数据库连接失败。", err)
		}
	case "sqlite":
		var err error
		dsn := viper.GetString("db.sqliteUrl")
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

}
