package survey

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/surveyApi"
	"github.com/sxz799/surveyX/middleware"
)

func Survey(e *gin.Engine) {
	g1 := e.Group("/api/admin/survey", middleware.JWTAuth())
	{
		g1.GET("/list", surveyApi.List)
		g1.POST("/", surveyApi.Add)
		g1.PUT("/", surveyApi.Update)
		g1.GET("/:id", surveyApi.Get)
		g1.DELETE("/:id", surveyApi.Del)
		g1.POST("/import", surveyApi.Import)
	}
	g2 := e.Group("/api/user/survey")
	{
		g2.GET("/list", surveyApi.List)
		g2.GET("/:id", surveyApi.Get)
	}
}
