package surveyApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/request"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
	"strconv"
)

var ss service.SurveyService

func List(c *gin.Context) {
	var pi request.PageInfo
	err := c.ShouldBindQuery(&pi)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if list, err := ss.List(pi); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}
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

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if result, err := ss.Get(id); err == nil {
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}
func Del(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = ss.Del(id); err == nil {
		response.OkWithMessage("删除成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}
