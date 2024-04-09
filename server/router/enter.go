package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/router/answer"
	"github.com/sxz799/surveyX/router/common"
	"github.com/sxz799/surveyX/router/question"
	"github.com/sxz799/surveyX/router/survey"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func RegRouter(e *gin.Engine) {
	survey.Survey(e)
	question.Question(e)
	answer.Answer(e)
	common.Common(e)
}

func RegWebRouter(e *gin.Engine, content embed.FS) {
	temp := template.Must(template.New("").ParseFS(content, "dist/index.html"))
	e.SetHTMLTemplate(temp)
	distFS, _ := fs.Sub(content, "dist")
	e.StaticFS("/dist", http.FS(distFS))
	e.NoRoute(func(context *gin.Context) {
		context.HTML(200, "index.html", "")
	})
	log.Println("已开启前后端整合模式！")
}
