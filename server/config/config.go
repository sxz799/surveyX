package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	ServerPort = "8080"
	GinMode    = "release"
)

func init() {
	log.Println("正在应用配置文件...")
	if _, err := os.Stat("conf.yaml"); os.IsNotExist(err) {
		log.Panicln("没找到配置文件...")
	} else {
		viper.SetConfigFile("conf.yaml")
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln("配置文件读取失败...")
	}

	ServerPort = viper.GetString("server.port")
	GinMode = viper.GetString("server.ginMode")
}
