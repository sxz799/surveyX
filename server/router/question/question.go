package question

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/question"
	"github.com/sxz799/surveyX/middleware"
)

func Question(e *gin.Engine) {
	g1 := e.Group("/api/admin/question", middleware.JWTAuth())
	{
		g1.GET("/list", question.List)
		g1.POST("/", question.Add)
		g1.PUT("/", question.Update)
		g1.GET("/:id", question.Get)
		g1.DELETE("/:id", question.Del)
		g1.GET("/analysis/:id", question.Analysis)

	}
	g2 := e.Group("/api/user/question")
	{

		g2.GET("/list", question.List)
	}
}
