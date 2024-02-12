package commonApi

import (
	"fmt"
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

	token, err := middleware.GenToken(u.Username)
	if err != nil {
		response.FailWithMessage("生成Token错误!", c)
		return
	}
	c.SetCookie("userId", fmt.Sprintf("%d", u.Id), 60*30, "", "", false, true)
	c.SetCookie("nickName", u.Nickname, 60*30, "", "", false, true)
	c.SetCookie("token", token, 60*30, "", "", false, true)
	response.OkWithDetailed(token, "登录成功", c)

}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", 0, "", "", false, true)
	c.SetCookie("nickName", "", 0, "", "", false, true)
	c.SetCookie("userId", "", 0, "", "", false, true)
	response.OkWithMessage("退出成功", c)
}
