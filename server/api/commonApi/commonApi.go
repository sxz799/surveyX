package commonApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	githubOauth2 "github.com/sxz799/surveyX/oauth/github"
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
	response.OkWithDetail(token, "登录成功", c)

}

func LoginByGithub(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("code为空", c)
		return
	}
	githubUser, err := githubOauth2.GetGithubUserInfo(code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	githubId := githubUser.ID
	user, _ := us.GetByGithubId(githubId)
	extMsg := ""
	if user.Id == 0 {
		user.Username = githubUser.Login
		user.Nickname = githubUser.Login
		user.Password = "123456"
		user.GithubUID = githubId
		us.Add(user)
		extMsg = "(已为您注册账号,账号:" + user.Username + ",密码:123456)"
	}
	token, err := middleware.GenToken(user.Id, user.Username, user.Nickname)
	if err != nil {
		response.FailWithMessage("生成Token错误", c)
		return
	}
	c.SetCookie("token", token, 60*30, "", "", false, true)
	response.OkWithDetail(token, "登录成功"+extMsg, c)

}

func Logout(c *gin.Context) {
	c.SetCookie("userId", "", 0, "", "", false, true)
	response.OkWithMessage("退出成功", c)
}
