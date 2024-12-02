package answer

import (
	"github.com/gin-gonic/gin"
	answerApi "github.com/sxz799/surveyX/api/answer"
	"github.com/sxz799/surveyX/middleware"
	answerService "github.com/sxz799/surveyX/service/answer"
)

func Answer(e *gin.Engine, service *answerService.Service) {
	api := answerApi.NewApi(service)

	g := e.Group("/api/user/answer")
	{
		g.POST("/", api.Add)
	}

	g2 := e.Group("/api/admin/answer", middleware.JWTAuth())
	{
		g2.GET("/list", api.List)
	}
}
