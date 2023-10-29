package question

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/questionApi"
)

func Question(e *gin.Engine) {
	g := e.Group("/api/question")
	{
		g.GET("/list/:id", questionApi.List)
		g.POST("/", questionApi.Add)
		g.PUT("/:id", questionApi.Update)
		g.GET("/:id", questionApi.Get)
		g.DELETE("/:id", questionApi.Del)
	}
}
