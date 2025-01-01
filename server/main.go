package main

import (
	"embed"
	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/router"
)

//go:embed dist
var content embed.FS

func main() {
	server, err := InitializeServer()
	if err != nil {
		log.Panicln("初始化服务失败:", err)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	router.RegRouter(r, server.DB, server.Config)
	router.RegWebRouter(r, content)

	log.Println("服务启动中,当前使用端口：", server.Config.Port)
	if err := r.Run(":" + server.Config.Port); err != nil {
		log.Panicln("服务启动失败。", err)
	}
}
