package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/router"
	"github.com/sxz799/surveyX/utils"
	"log"
)

//go:embed dist
var content embed.FS

func main() {
	config.Init()

	startGin()
}

func startGin() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	db := utils.InitDB()
	router.RegRouter(r, db)
	router.RegWebRouter(r, content)
	log.Println("服务启动中,当前使用端口：", config.Port)

	err := r.Run(":" + config.Port)
	if err != nil {
		log.Panicln("服务启动失败。", err)
	}
}
