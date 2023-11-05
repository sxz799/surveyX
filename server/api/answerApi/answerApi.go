package answerApi

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service"
)

var as service.AnswerService

// Add godoc
// @Summary 添加
// @Description 添加答案
// @Tags 答案
// @Accept json
// @Produce json
// @Param answer body []entity.Answer true "答案"
// @Router /answer/ [post]
func Add(c *gin.Context) {
	var a []entity.Answer
	err := c.ShouldBind(&a)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = as.Add(a); err == nil {
		response.OkWithMessage("提交成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}
