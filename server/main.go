package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/router"
	"github.com/sxz799/surveyX/utils"
	"log"
)

func main() {
	utils.InitDB()
	gin.SetMode(config.GinMode)
	r := gin.Default()
	r.Use(cors.Default())
	router.RegRouter(r)
	log.Println("服务启动中,当前使用端口：", config.ServerPort)
	_ = r.Run(":" + config.ServerPort)
}
