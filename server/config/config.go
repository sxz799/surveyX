package config

import (
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config 存储所有配置信息
type Config struct {
	Port              string
	SqlType           string
	SqlUrl            string
	OauthClientId     string
	OauthClientSecret string
}

// NewConfig 创建并初始化配置
func NewConfig() *Config {
	log.Println("正在读取配置文件...")
	cfg := &Config{}

	// 设置默认值
	cfg.Port = "65534"
	cfg.SqlType = "sqlite"
	cfg.SqlUrl = "survey.db"
	cfg.OauthClientId = "ff3f2aa615877bd961e7"
	cfg.OauthClientSecret = "84f840226a1fe34990e4725d9040fa9e6086da36"

	// 读取配置文件
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	configName := "config-" + env
	viper.SetConfigName(configName)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("没找到配置文件,使用默认配置.err:", err.Error())
	} else {
		log.Println("使用中的配置:", configName)
		cfg.Port = viper.GetString("port")
		cfg.SqlType = viper.GetString("db.sqlType")
		cfg.SqlUrl = viper.GetString("db.sqlUrl")
		cfg.OauthClientId = viper.GetString("oauth.clientId")
		cfg.OauthClientSecret = viper.GetString("oauth.clientSecret")
	}

	// 读取命令行参数
	port := flag.String("port", cfg.Port, "端口号")
	sqlType := flag.String("sqlType", cfg.SqlType, "数据库类型")
	sqlUrl := flag.String("sqlUrl", cfg.SqlUrl, "数据库地址")
	clientId := flag.String("clientId", cfg.OauthClientId, "clientId")
	clientSecret := flag.String("clientSecret", cfg.OauthClientSecret, "clientSecret")
	flag.Parse()

	cfg.Port = *port
	cfg.SqlType = *sqlType
	cfg.SqlUrl = *sqlUrl
	cfg.OauthClientId = *clientId
	cfg.OauthClientSecret = *clientSecret

	return cfg
}
