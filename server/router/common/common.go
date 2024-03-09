package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/commonApi"
	githubOauth2 "github.com/sxz799/surveyX/oauth/github"
)

func Common(e *gin.Engine) {

	e.POST("/api/login", commonApi.Login)
	e.POST("/api/logout", commonApi.Logout)
	e.POST("/api/oauth/github", commonApi.LoginByGithub)
	e.GET("/api/oauth/loginUrl", githubOauth2.GetLoginUrl)

}
