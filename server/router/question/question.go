package question

import (
	"github.com/gin-gonic/gin"
	questionApi "github.com/sxz799/surveyX/api/question"
	"github.com/sxz799/surveyX/middleware"
	questionService "github.com/sxz799/surveyX/service/question"
)

func Question(e *gin.Engine, service *questionService.Service) {
	api := questionApi.NewApi(service)

	g1 := e.Group("/api/admin/question", middleware.JWTAuth())
	{
		g1.GET("/list", api.List)
		g1.POST("", api.Add)
		g1.PUT("", api.Update)
		g1.GET("/:id", api.Get)
		g1.DELETE("/:id", api.Del)
		g1.GET("/analysis/:id", api.Analysis)
	}

	g2 := e.Group("/api/guest/question")
	{
		g2.GET("/list", api.List)
	}
}
