package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service/user"
)

type Api struct {
	userService *user.Service
}

func NewApi(us *user.Service) *Api {
	return &Api{
		userService: us,
	}
}

func (a *Api) Register(c *gin.Context) {
	var u entity.User
	err := c.ShouldBind(&u)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}

	if _, err = a.userService.Add(u); err == nil {
		response.OkWithMessage("注册成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

func (a *Api) ChangePwd(c *gin.Context) {
	var u entity.User
	err := c.ShouldBind(&u)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = a.userService.Update(u); err == nil {
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

func (a *Api) ResetPwd(c *gin.Context) {
	var tu entity.User
	err := c.ShouldBind(&tu)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}

	if u, err := a.userService.GetByUsernameAndEmail(tu.Username, tu.Email); err == nil {
		u.Password = tu.Password
		err = a.userService.Update(u)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}
		response.OkWithMessage("密码修改成功！", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}
