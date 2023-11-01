package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/router/answer"
	"github.com/sxz799/surveyX/router/question"
	"github.com/sxz799/surveyX/router/survey"
	"log"
	"os"
)

func RegRouter(e *gin.Engine) {
	survey.Survey(e)
	question.Question(e)
	answer.Answer(e)

	_, err := os.Stat("dist")
	if err == nil {
		e.LoadHTMLGlob("dist/index.html")
		e.Static("/dist", "dist")
		e.GET("/", func(context *gin.Context) {
			context.HTML(200, "index.html", "")
		})
		e.GET("/admin", func(context *gin.Context) {
			context.HTML(200, "index.html", "")
		})
		log.Println("已开启前后端整合模式！")
	}
}
