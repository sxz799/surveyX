package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/router/answer"
	"github.com/sxz799/surveyX/router/common"
	"github.com/sxz799/surveyX/router/question"
	"github.com/sxz799/surveyX/router/survey"
	"github.com/sxz799/surveyX/router/user"
	"github.com/sxz799/surveyX/service"
	"gorm.io/gorm"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func RegRouter(e *gin.Engine, db *gorm.DB, config *config.Config) {
	e.Use(cors())
	// 创建服务容器
	container := service.NewContainer(db)

	// 注册各模块路由
	survey.Survey(e, container.SurveyService)
	question.Question(e, container.QuestionService)
	answer.Answer(e, container.AnswerService)
	common.Common(e, container.CommonService, container.UserService, config)
	user.User(e, container.UserService)
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
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
