package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/docs"
	"github.com/sxz799/surveyX/router"
	"github.com/sxz799/surveyX/utils"
	"log"
)

// @title 问卷调查SurveyX API
// @version 1.0
func main() {
	utils.InitDB()
	gin.SetMode(config.GinMode)
	r := gin.Default()
	r.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/api"
	router.RegRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("服务启动中,当前使用端口：", config.ServerPort)
	_ = r.Run(":" + config.ServerPort)
}
