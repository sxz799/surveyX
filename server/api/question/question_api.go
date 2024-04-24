package question

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
	"strconv"
)

var qs service.QuestionService

// List godoc
func List(c *gin.Context) {
	var q entity.QuestionSearch
	err := c.ShouldBindQuery(&q)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if list, err := qs.List(q); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}

// Add godoc
func Add(c *gin.Context) {
	var q entity.Question
	err := c.ShouldBind(&q)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = qs.Add(q); err == nil {
		response.OkWithMessage("添加成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Update godoc
func Update(c *gin.Context) {
	var q entity.Question
	err := c.ShouldBind(&q)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = qs.Update(q); err == nil {
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// Get godoc
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if result, err := qs.Get(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Del godoc
func Del(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = qs.Del(id); err == nil {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Analysis godoc
func Analysis(c *gin.Context) {
	id := c.Param("id")
	if result, err := qs.Analysis(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}
