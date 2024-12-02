package question

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service/question"
	"strconv"
)

type Api struct {
	questionService *question.Service
}

func NewApi(qs *question.Service) *Api {
	return &Api{
		questionService: qs,
	}
}

// List godoc
func (a *Api) List(c *gin.Context) {
	var q entity.QuestionSearch
	err := c.ShouldBindQuery(&q)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if list, err := a.questionService.List(q); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}

// Add godoc
func (a *Api) Add(c *gin.Context) {
	var q entity.Question
	err := c.ShouldBind(&q)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = a.questionService.Add(q); err == nil {
		response.OkWithMessage("添加成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Update godoc
func (a *Api) Update(c *gin.Context) {
	var q entity.Question
	err := c.ShouldBind(&q)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = a.questionService.Update(q); err == nil {
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// Get godoc
func (a *Api) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if result, err := a.questionService.Get(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Del godoc
func (a *Api) Del(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = a.questionService.Del(id); err == nil {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Analysis godoc
func (a *Api) Analysis(c *gin.Context) {
	id := c.Param("id")
	if result, err := a.questionService.Analysis(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}
