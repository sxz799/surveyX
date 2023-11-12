package answer

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/answerApi"
	"github.com/sxz799/surveyX/middleware"
)

func Answer(e *gin.Engine) {
	g := e.Group("/api/user/answer")
	{
		g.POST("/", answerApi.Add)
	}

	g2 := e.Group("/api/admin/answer", middleware.JWTAuth())
	{
		g2.GET("/list", answerApi.List)
	}
}
