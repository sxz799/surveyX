package commonApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/middleware"
	"github.com/sxz799/surveyX/model/common/response"
)

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
	if k.Key != "123456" {
		response.FailWithMessage("Key错误!", c)
		return
	}
	token, err := middleware.GenToken(k.Key)
	if err != nil {
		response.FailWithMessage("生成Token错误!", c)
		return
	}
	c.SetCookie("token", token, 3600, "", "", false, true)
	response.OkWithDetailed(token, "登录成功", c)

}
