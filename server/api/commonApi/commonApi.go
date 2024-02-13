package commonApi

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
)

var us service.UserService

func GetCurrentUser(c *gin.Context) (userInfo entity.LoginUser) {
	session := sessions.Default(c)
	userInfo = session.Get("currentUser").(entity.LoginUser) // 类型转换一下
	return
}

func setCurrentUser(c *gin.Context, userInfo entity.LoginUser) {
	session := sessions.Default(c)
	session.Set("currentUser", userInfo)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	session.Save()
}

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
	var userInfo entity.LoginUser
	userInfo.Id = u.Id
	userInfo.Username = u.Username
	userInfo.Nickname = u.Nickname
	userInfo.Email = u.Email
	userInfo.Phone = u.Phone
	setCurrentUser(c, userInfo)
	c.SetCookie("token", token, 60*30, "", "", false, true)
	response.OkWithDetailed(token, "登录成功", c)

}

func Logout(c *gin.Context) {
	c.SetCookie("userId", "", 0, "", "", false, true)
	response.OkWithMessage("退出成功", c)
}
