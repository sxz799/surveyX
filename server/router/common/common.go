package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/commonApi"
)

func Common(e *gin.Engine) {

	e.POST("/api/login", commonApi.Login)
	e.POST("/api/logout", commonApi.Logout)

}
