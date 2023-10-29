package survey

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/surveyApi"
)

func Survey(e *gin.Engine) {
	g := e.Group("/api/survey")
	{
		g.GET("/list", surveyApi.List)
		g.POST("/", surveyApi.Add)
		g.PUT("/:id", surveyApi.Update)
		g.GET("/:id", surveyApi.Get)
		g.DELETE("/:id", surveyApi.Del)
	}
}
