package commonApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
)

var us service.UserService

func Login(c *gin.Context) {

	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		response.FailWithMessage("参数错误!", c)
		return
	}
	u, err := us.Login(user)
	if err != nil {
		response.FailWithMessage("账号或密码错误!", c)
		return
	}

	token, err := middleware.GenToken(u.Id, u.Username, u.Nickname)
	if err != nil {
		response.FailWithMessage("生成Token错误!", c)
		return
	}
	c.SetCookie("token", token, 60*30, "", "", false, true)
	response.OkWithDetailed(token, "登录成功", c)

}

func Logout(c *gin.Context) {
	c.SetCookie("userId", "", 0, "", "", false, true)
	response.OkWithMessage("退出成功", c)
}
