package router

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/router/answer"
	"github.com/sxz799/surveyX/router/common"
	"github.com/sxz799/surveyX/router/question"
	"github.com/sxz799/surveyX/router/survey"
	"github.com/sxz799/surveyX/service"
	"gorm.io/gorm"
)

func RegRouter(e *gin.Engine, db *gorm.DB) {
	// 创建服务容器
	container := service.NewContainer(db)

	// 注册各模块路由
	survey.Survey(e, container.SurveyService)
	question.Question(e, container.QuestionService)
	answer.Answer(e, container.AnswerService)
	common.Common(e, container.UserService)
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
