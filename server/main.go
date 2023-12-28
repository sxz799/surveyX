package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/router"
	"github.com/sxz799/surveyX/utils"
)

// ------前后端分离调试时请注释下面代码------
//
//go:embed dist
var content embed.FS

//------前后端分离调试时请注释上面代码------

func main() {
	config.Init()
	utils.InitDB()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	
	//------前后端分离调试时请注释下面代码------
	temp := template.Must(template.New("").ParseFS(content, "dist/index.html"))
	r.SetHTMLTemplate(temp)
	distFS, _ := fs.Sub(content, "dist")
	r.StaticFS("/dist", http.FS(distFS))
	r.NoRoute(func(context *gin.Context) {
		context.HTML(200, "index.html", "")
	})
	log.Println("已开启前后端整合模式！")
	//------前后端分离调试时请注释上面代码------


	router.RegRouter(r)
	log.Println("服务启动中,当前使用端口：", config.Port)

	err := r.Run(":" + config.Port)
	if err != nil {
		log.Panicln("服务启动失败。", err)
	}
}
