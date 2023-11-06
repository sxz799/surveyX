package commonApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/model/common/response"
)

// Login godoc
// @Summary 登录
// @Description 登录
// @Tags 公共
// @Accept json
// @Produce json
// @Param key body string true "key"
// @Router /login [post]
func Login(c *gin.Context) {
	type Key struct {
		Key string `json:"key" form:"key" binding:"required"`
	}
	var k Key
	err := c.ShouldBind(&k)
	if err != nil {
		response.FailWithMessage("参数错误!", c)
		return
	}
	if k.Key != config.Key {
		response.FailWithMessage("Key错误!", c)
		return
	}
	token, err := middleware.GenToken(k.Key)
	if err != nil {
		response.FailWithMessage("生成Token错误!", c)
		return
	}
	c.SetCookie("token", token, 60*30, "", "", false, true)
	response.OkWithDetailed(token, "登录成功", c)

}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", 0, "", "", false, true)
	response.OkWithMessage("退出成功", c)
}
