package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	githubOauth2 "github.com/sxz799/surveyX/oauth/github"
	"github.com/sxz799/surveyX/service/common"
	"github.com/sxz799/surveyX/service/user"
	"log"
)

type Api struct {
	commonService *common.Service
	userService   *user.Service
	githubOAuth   *githubOauth2.GithubOAuth
}

func NewApi(cs *common.Service, us *user.Service, githubOAuth *githubOauth2.GithubOAuth) *Api {
	return &Api{
		userService:   us,
		commonService: cs,
		githubOAuth:   githubOAuth,
	}
}

func (a *Api) Login(c *gin.Context) {

	var user entity.LoginUser
	err := c.ShouldBind(&user)
	if err != nil {
		response.FailWithMessage("参数错误!", c)
		return
	}
	u, err := a.commonService.Login(user)
	if err != nil {
		response.FailWithMessage("账号或密码错误!", c)
		return
	}

	token, err := middleware.GenToken(u, user.RememberMe)
	if err != nil {
		response.FailWithMessage("生成Token错误!", c)
		return
	}
	u.Password = ""
	response.OkWithDetail(gin.H{"token": token, "userInfo": u}, "登录成功", c)

}

func (a *Api) LoginByGithub(c *gin.Context) {
	code := c.Query("code")
	log.Println("code:", code)
	if code == "" {
		response.FailWithMessage("code为空", c)
		return
	}
	githubUser, err := a.githubOAuth.GetGithubUserInfo(code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	githubId := githubUser.ID
	u, _ := a.userService.GetByGithubId(githubId)
	extMsg := ""
	if u.Id == 0 {
		u.Username = githubUser.Login
		u.Nickname = githubUser.Login
		u.Password = "123456"
		u.GithubUID = githubId
		id, _ := a.userService.Add(u)
		u.Id = id
		extMsg = "(已为您注册账号,账号:" + u.Username + ",密码:123456)"
	}
	token, err := middleware.GenToken(u, false)
	if err != nil {
		response.FailWithMessage("生成Token错误", c)
		return
	}
	u.Password = ""
	response.OkWithDetail(gin.H{"token": token, "userInfo": u}, "登录成功"+extMsg, c)

}

func (a *Api) Logout(c *gin.Context) {
	response.OkWithMessage("退出成功", c)
}
