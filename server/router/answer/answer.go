package answer

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/answerApi"
)

func Answer(e *gin.Engine) {
	g := e.Group("/api/answer")
	{
		g.POST("/", answerApi.Add)
	}
}
