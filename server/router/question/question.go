package question

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/questionApi"
	"github.com/sxz799/surveyX/middleware"
)

func Question(e *gin.Engine) {
	g1 := e.Group("/api/admin/question", middleware.JWTAuth())
	{
		g1.GET("/list", questionApi.List)
		g1.POST("/", questionApi.Add)
		g1.PUT("/", questionApi.Update)
		g1.GET("/:id", questionApi.Get)
		g1.DELETE("/:id", questionApi.Del)
		g1.GET("/analysis/:id", questionApi.Analysis)

	}
	g2 := e.Group("/api/user/question")
	{

		g2.GET("/list", questionApi.List)
	}
}
