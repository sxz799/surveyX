package commonApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	githubOauth2 "github.com/sxz799/surveyX/oauth/github"
	"github.com/sxz799/surveyX/service"
	"log"
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
	u.Password = ""
	response.OkWithDetail(gin.H{"token": token, "userInfo": u}, "登录成功", c)

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
	u, _ := us.GetByGithubId(githubId)
	extMsg := ""
	if u.Id == 0 {
		u.Username = githubUser.Login
		u.Nickname = githubUser.Login
		u.Password = "123456"
		u.GithubUID = githubId
		id, _ := us.Add(u)
		u.Id = id
		extMsg = "(已为您注册账号,账号:" + u.Username + ",密码:123456)"
	}
	token, err := middleware.GenToken(u.Id, u.Username, u.Nickname)
	if err != nil {
		response.FailWithMessage("生成Token错误", c)
		return
	}
	u.Password = ""
	response.OkWithDetail(gin.H{"token": token, "userInfo": u}, "登录成功"+extMsg, c)

}

func ChangePwd(c *gin.Context) {
	var u entity.User
	err := c.ShouldBind(&u)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	log.Println(u)
	if err = us.Update(u); err == nil {
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

func Logout(c *gin.Context) {
	response.OkWithMessage("退出成功", c)
}
