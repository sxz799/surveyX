package answer

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/answer"
	"github.com/sxz799/surveyX/middleware"
)

func Answer(e *gin.Engine) {
	g := e.Group("/api/user/answer")
	{
		g.POST("/", answer.Add)
	}

	g2 := e.Group("/api/admin/answer", middleware.JWTAuth())
	{
		g2.GET("/list", answer.List)
	}
}
