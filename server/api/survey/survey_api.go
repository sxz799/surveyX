package survey

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service/survery"
)

type Api struct {
	surveyService *survery.Service
}

func NewApi(ss *survery.Service) *Api {
	return &Api{
		surveyService: ss,
	}
}

// List godoc
func (a *Api) List(c *gin.Context) {
	var s entity.SurveySearch
	err := c.ShouldBindQuery(&s)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}

	value, exists := c.Get("userInfo")
	if !exists {
		response.FailWithExpire(c)
		c.Abort()
		return
	}
	s.Survey.UserId = value.(entity.User).Id

	if list, err := a.surveyService.List(s); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}

// Add godoc
func (a *Api) Add(c *gin.Context) {
	var s entity.Survey
	err := c.ShouldBind(&s)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	value, exists := c.Get("userInfo")
	if !exists {
		response.FailWithExpire(c)
		c.Abort()
		return
	}
	s.UserId = value.(entity.User).Id
	if err = a.surveyService.Add(s); err == nil {
		response.OkWithMessage("添加成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Update godoc
func (a *Api) Update(c *gin.Context) {
	var s entity.Survey
	err := c.ShouldBind(&s)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = a.surveyService.Update(s); err == nil {
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// Get godoc
func (a *Api) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数有误", c)
		return
	}
	if result, err := a.surveyService.Get(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Del godoc
func (a *Api) Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err := a.surveyService.Del(id); err == nil {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// Import godoc
func (a *Api) Import(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("没有获取到文件!", c)
		return
	}
	value, exists := c.Get("userInfo")
	if !exists {
		response.FailWithExpire(c)
		c.Abort()
		return
	}
	err = a.surveyService.Import(value.(entity.User).Id, file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("导入成功", c)
	}
}

// Analysis godoc
func (a *Api) Analysis(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数有误", c)
		return
	}
	if result, err := a.surveyService.Analysis(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}
