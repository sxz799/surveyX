package user

import (
	"github.com/gin-gonic/gin"
	userApi "github.com/sxz799/surveyX/api/user"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/service/user"
)

func User(e *gin.Engine, us *user.Service) {

	api := userApi.NewApi(us)
	g := e.Group("/api/user")
	{
		g.POST("/register", api.Register)
		g.PUT("/changePwd", middleware.JWTAuth(), api.ChangePwd)
		g.PUT("/resetPwd", api.ResetPwd)
	}

}
