package questionApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
	"strconv"
)

var qs service.QuestionService

// List godoc
// @Summary 题目列表
// @Description 获取题目列表列表
// @Tags 题目
// @Produce  json
// @Param page query request.PageInfo false "分页信息"
// @Param id path int true "题目id"
// @Router /question/list [get]
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
// @Summary 添加
// @Description 添加题目
// @Tags 题目
// @Accept json
// @Produce json
// @Param question body entity.Question true "题目"
// @Router /question/ [post]
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
// @Description 更新题目
// @Tags 题目
// @Accept json
// @Produce json
// @Param question body entity.Question true "题目"
// @Router /question/ [put]
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
// @Summary 获取
// @Description 获取题目
// @Tags 题目
// @Produce  json
// @Router /question/{id} [get]
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
// @Description 删除题目
// @Tags 题目
// @Produce  json
// @Router /question/{id} [delete]
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
