package surveyApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
)

var ss service.SurveyService

// List godoc
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

// Import godoc
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
