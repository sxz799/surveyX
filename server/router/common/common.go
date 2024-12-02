package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/common"
	"github.com/sxz799/surveyX/middleware"
	githubOauth2 "github.com/sxz799/surveyX/oauth/github"
	"github.com/sxz799/surveyX/service/user"
)

func Common(e *gin.Engine, service *user.Service) {
	api := common.NewApi(service)
	e.POST("/api/login", api.Login)
	e.POST("/api/logout", api.Logout)
	e.POST("/api/oauth/github", api.LoginByGithub)
	e.GET("/api/oauth/loginUrl", githubOauth2.GetLoginUrl)
	e.PUT("/api/changePwd", middleware.JWTAuth(), api.ChangePwd)

}
