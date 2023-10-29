package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/router/question"
	"github.com/sxz799/surveyX/router/survey"
)

func RegRouter(e *gin.Engine) {
	survey.Survey(e)
	question.Question(e)
}
