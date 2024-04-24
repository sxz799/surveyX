package survey

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/survey"
	"github.com/sxz799/surveyX/middleware"
)

func Survey(e *gin.Engine) {
	g1 := e.Group("/api/admin/survey", middleware.JWTAuth())
	{
		g1.GET("/list", survey.List)
		g1.POST("/", survey.Add)
		g1.PUT("/", survey.Update)
		g1.GET("/:id", survey.Get)
		g1.DELETE("/:id", survey.Del)
		g1.POST("/import", survey.Import)
		g1.GET("/analysis/:id", survey.Analysis)
	}
	g2 := e.Group("/api/user/survey")
	{
		g2.GET("/:id", survey.Get)
	}
}
