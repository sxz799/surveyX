package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	ServerPort  string
	GinMode     string
	Key         string
	SqlType     string
	SqliteUrl   string
	MysqlUrl    string
	PostgresUrl string
)

func init() {
	log.Println("正在应用配置文件...")
	if _, err := os.Stat("conf.yaml"); os.IsNotExist(err) {
		log.Println("没找到配置文件,使用默认配置...")
		ServerPort = "3000"
		GinMode = "release"
		Key = "123456"
		SqlType = "sqlite"
		SqliteUrl = "surveyX.db"
	} else {
		viper.SetConfigFile("conf.yaml")
		err = viper.ReadInConfig()
		if err != nil {
			log.Panicln("配置文件读取失败...")
		}
		ServerPort = viper.GetString("server.port")
		GinMode = viper.GetString("server.ginMode")
		Key = viper.GetString("admin.key")
		SqlType = viper.GetString("db.sqlType")
		SqliteUrl = viper.GetString("db.sqliteUrl")
		MysqlUrl = viper.GetString("db.mysqlUrl")
		PostgresUrl = viper.GetString("db.postgresUrl")
	}

}
