package answer

import (
	"github.com/gin-gonic/gin"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service/answer"
)

type Api struct {
	answerService *answer.Service
}

func NewApi(as *answer.Service) *Api {
	return &Api{
		answerService: as,
	}
}

// Add godoc
func (a *Api) Add(c *gin.Context) {
	var answers []entity.Answer
	err := c.ShouldBind(&answers)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if err = a.answerService.Add(answers); err == nil {
		response.OkWithMessage("提交成功", c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}

}

// List godoc
func (a *Api) List(c *gin.Context) {
	var as entity.AnswerSearch
	err := c.ShouldBindQuery(&as)
	if err != nil {
		response.FailWithMessage("参数有误", c)
		return
	}
	if list, err := a.answerService.List(as); err == nil {
		response.OkWithData(list, c)
	} else {
		response.Fail(c)
	}
}
