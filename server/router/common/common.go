package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/api/commonApi"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/model/common/response"
)

func Common(e *gin.Engine) {

	e.POST("/api/login", commonApi.Login)
	e.POST("/api/logout", commonApi.Logout)
	e.POST("/api/oauth/github", commonApi.LoginByGithub)
	e.GET("/api/oauth/loginUrl", func(c *gin.Context) {
		response.OkWithData("https://github.com/login/oauth/authorize?client_id="+config.OauthClientId, c)
	})

}
