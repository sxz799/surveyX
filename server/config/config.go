package config

import (
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	Port              string
	SqlType           string
	SqlUrl            string
	OauthClientId     string
	OauthClientSecret string
)

func Init() {
	log.Println("正在读取配置文件...")
	viper.SetConfigType("yaml") // 配置文件类型
	viper.AddConfigPath(".")    // 配置文件路径
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // 默认环境
	}
	configName := "config-" + env
	viper.SetConfigName(configName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("没找到配置文件,使用默认配置.err:", err.Error())
		Port = "65534"
		SqlType = "sqlite"
		SqlUrl = "survey.db"
		OauthClientId = "ff3f2aa615877bd961e7"
		OauthClientSecret = "84f840226a1fe34990e4725d9040fa9e6086da36"
	} else {
		log.Println("使用中的配置:", configName)
		Port = viper.GetString("port")
		SqlType = viper.GetString("db.sqlType")
		SqlUrl = viper.GetString("db.sqlUrl")
		OauthClientId = viper.GetString("oauth.clientId")
		OauthClientSecret = viper.GetString("oauth.clientSecret")
	}
	// 读取命令行参数
	port := flag.String("port", Port, "端口号")
	sqlType := flag.String("sqlType", SqlType, "数据库类型")
	sqlUrl := flag.String("sqlUrl", SqlUrl, "数据库地址")
	clientId := flag.String("clientId", OauthClientId, "clientId")
	clientSecret := flag.String("clientSecret", OauthClientSecret, "clientSecret")
	flag.Parse()
	Port = *port
	SqlType = *sqlType
	SqlUrl = *sqlUrl
	OauthClientId = *clientId
	OauthClientSecret = *clientSecret
}
