package surveyApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
)

var ss service.SurveyService

// List godoc
// @Summary 问卷列表
// @Description 获取问卷列表
// @Tags 问卷
// @Produce  json
// @Param page query request.PageInfo false "分页信息"
// @Router /survey/list [get]
func List(c *gin.Context) {
	var s entity.SurveySearch
	err := c.ShouldBindQuery(&s)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if list, err := ss.List(s); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}

// Add godoc
// @Summary 添加
// @Description 添加问卷
// @Tags 问卷
// @Accept json
// @Produce json
// @Param survey body entity.Survey true "问卷"
// @Router /survey/ [post]
func Add(c *gin.Context) {
	var s entity.Survey
	err := c.ShouldBind(&s)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = ss.Add(s); err == nil {
		response.OkWithMessage("添加成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Update godoc
// @Summary 更新
// @Description 更新问卷
// @Tags 问卷
// @Accept json
// @Produce json
// @Param survey body entity.Survey true "问卷"
// @Router /survey/ [put]
func Update(c *gin.Context) {
	var s entity.Survey
	err := c.ShouldBind(&s)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = ss.Update(s); err == nil {
		response.OkWithMessage("更新成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// Get godoc
// @Summary 获取
// @Description 获取获取
// @Tags 问卷
// @Accept json
// @Produce json
// @Param id path int true "题目id"
// @Router /survey/{id} [get]
func Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数有误", c)
		return
	}
	if result, err := ss.Get(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// Del godoc
// @Summary 删除
// @Description 删除问卷
// @Tags 问卷
// @Param id path int true "问卷id"
// @Router /survey/{id} [delete]
func Del(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err := ss.Del(id); err == nil {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

func Import(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("没有获取到文件!", c)
		return
	}

	err = ss.Import(file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("导入成功", c)
	}
}
