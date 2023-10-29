package questionApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/request"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
	"strconv"
)

var qs service.QuestionService

// @BasePath /api

// List godoc
// @Summary 问卷调查列表
// @Description 获取问卷调查列表
// @Schemes
// @Tags 问卷调查
// @Router /survey/list [get]
func List(c *gin.Context) {
	var pi request.PageInfo
	err := c.ShouldBindQuery(&pi)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if list, err := qs.List(pi); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}

// Add godoc
// @Summary 新建
// @Description 新建问卷调查
// @Tags 问卷调查
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
// @Summary 更新
// @Description 更新问卷调查
// @Tags 问卷调查
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
// @Summary 查询
// @Description 获取一个问卷调查
// @Tags 问卷调查
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
// @Summary 删除
// @Description 删除一个问卷调查
// @Tags 问卷调查
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
