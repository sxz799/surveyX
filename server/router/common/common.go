package common

import (
	"github.com/gin-gonic/gin"
	commonApi "github.com/sxz799/surveyX/api/common"
	"github.com/sxz799/surveyX/config"
	githubOauth2 "github.com/sxz799/surveyX/oauth/github"
	commonService "github.com/sxz799/surveyX/service/common"
	"github.com/sxz799/surveyX/service/user"
)

func Common(e *gin.Engine, cs *commonService.Service, us *user.Service, config *config.Config) {
	githubOAuth := githubOauth2.NewGithubOAuth(config.OauthClientId, config.OauthClientSecret)
	api := commonApi.NewApi(cs, us, githubOAuth)
	g := e.Group("/api/common")
	{
		g.POST("/login", api.Login)
		g.POST("/logout", api.Logout)
		g.POST("/oauth/github", api.LoginByGithub)
		g.GET("/oauth/loginUrl", githubOAuth.GetLoginUrl)
	}
}
