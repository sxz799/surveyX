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
	database := viper.GetString("db.database")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	switch sqlType {
	case "postgres":
		dsn := "postgres://" + username + ":" + password + "@" + host + "/" + database
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("postgres数据库连接失败。", err)
		}
	case "mysql":
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("mysql数据库连接失败。", err)
		}
	case "sqlite":
		var err error
		DB, err = gorm.Open(sqlite.Open(database+".db"), &gorm.Config{})
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
