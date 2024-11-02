package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/common"
	"github.com/sxz799/surveyX/middleware"
	githubOauth2 "github.com/sxz799/surveyX/oauth/github"
)

func Common(e *gin.Engine) {

	e.POST("/api/login", common.Login)
	e.POST("/api/logout", common.Logout)
	e.POST("/api/oauth/github", common.LoginByGithub)
	e.GET("/api/oauth/loginUrl", githubOauth2.GetLoginUrl)
	e.PUT("/api/changePwd", middleware.JWTAuth(), common.ChangePwd)

}
