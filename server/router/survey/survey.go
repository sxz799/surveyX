package survey

import (
	"github.com/gin-gonic/gin"
	surveyApi "github.com/sxz799/surveyX/api/survey"
	"github.com/sxz799/surveyX/middleware"
	surveyService "github.com/sxz799/surveyX/service/survery"
)

func Survey(e *gin.Engine, service *surveyService.Service) {
	api := surveyApi.NewApi(service)

	g1 := e.Group("/api/admin/survey", middleware.JWTAuth())
	{
		g1.GET("/list", api.List)
		g1.POST("/", api.Add)
		g1.PUT("/", api.Update)
		g1.GET("/:id", api.Get)
		g1.DELETE("/:id", api.Del)
		g1.POST("/import", api.Import)
		g1.GET("/analysis/:id", api.Analysis)
	}

	g2 := e.Group("/api/guest/survey")
	{
		g2.GET("/:id", api.Get)
	}
}
