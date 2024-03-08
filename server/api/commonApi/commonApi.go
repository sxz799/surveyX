package commonApi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
	"github.com/sxz799/surveyX/utils"
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

func LogoutByGithub(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("code为空", c)
		return
	}
	header := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	body := map[string]string{
		"client_id":     config.OauthClientId,
		"client_secret": config.OauthClientSecret,
		"code":          code,
	}
	marshal, _ := json.Marshal(body)
	request, err := utils.SendHTTPRequest("POST", "https://github.com/login/oauth/access_token", header, marshal)
	if err != nil {
		response.FailWithMessage("请求Github错误", c)
		return
	}
	var githubToken map[string]string
	_ = json.Unmarshal(request, &githubToken)
	accessToken := githubToken["access_token"]
	if accessToken == "" {
		response.FailWithMessage("获取accessToken错误", c)
		return
	}
	header = map[string]string{
		"Authorization": "bearer " + accessToken,
		"Content-Type":  "application/json",
	}
	request, err = utils.SendHTTPRequest("GET", "https://api.github.com/user", header, nil)
	if err != nil {
		response.FailWithMessage("请求Github错误", c)
		return
	}
	var githubUser map[string]interface{}
	_ = json.Unmarshal(request, &githubUser)
	githubId := githubUser["id"].(float64)
	user, _ := us.GetByGithubId(githubId)
	extMsg := ""
	if user.Id == 0 {
		user.Username = githubUser["login"].(string)
		user.Nickname = githubUser["login"].(string)
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
	response.OkWithDetailed(token, "登录成功"+extMsg, c)

}

func Logout(c *gin.Context) {
	c.SetCookie("userId", "", 0, "", "", false, true)
	response.OkWithMessage("退出成功", c)
}
