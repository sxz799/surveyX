package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
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
	if _, err := os.Stat("conf.yaml"); os.IsNotExist(err) {
		log.Println("没找到配置文件,使用默认配置...")
		Port = "65534"
		SqlType = "sqlite"
		SqlUrl = "survey.db"
		OauthClientId = "ff3f2aa615877bd961e7"
		OauthClientSecret = "84f840226a1fe34990e4725d9040fa9e6086da36"
	} else {
		viper.SetConfigFile("conf.yaml")
		err = viper.ReadInConfig()
		if err != nil {
			log.Panicln("配置文件读取失败...")
		}
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
