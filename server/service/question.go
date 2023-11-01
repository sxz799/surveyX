package service

import (
	"github.com/sxz799/surveyX/model/common/request"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
)

type QuestionService struct {
}

var optionService OptionService

func (ts *QuestionService) List(pi request.PageInfo, sId int) (response.PageResult, error) {
	var qs []entity.Question
	var total int64
	limit := pi.PageSize
	offset := pi.PageSize * (pi.PageNum - 1)
	db := utils.DB.Model(&entity.Question{})
	db.Where("survey_id=?", sId)
	db.Count(&total)
	err := db.Debug().Limit(limit).Offset(offset).Order("`order`").Find(&qs).Error
	for i := range qs {
		ops := optionService.List(qs[i].Id)
		qs[i].Options = ops
	}

	return response.PageResult{
		List:     qs,
		Total:    total,
		Page:     pi.PageNum,
		PageSize: pi.PageSize}, err
}

func (ts *QuestionService) Add(q entity.Question) (err error) {

	err = utils.DB.Debug().Create(&q).Error

	for i := range q.Options {
		q.Options[i].QuestionId = q.Id
		q.Options[i].SurveyId = q.SurveyId
	}

	optionService.Add(q.Options)
	return
}

func (ts *QuestionService) Update(q entity.Question) (err error) {

	err = utils.DB.Debug().Updates(&q).Error
	for i := range q.Options {
		q.Options[i].QuestionId = q.Id
		q.Options[i].SurveyId = q.SurveyId
	}
	optionService.Del(q.Id)
	optionService.Add(q.Options)
	return
}

func (ts *QuestionService) Del(id int) (err error) {
	q := entity.Question{
		Id: id,
	}
	err = utils.DB.Delete(&q).Error
	optionService.Del(id)
	return
}

func (ts *QuestionService) DelBySurveyId(surveyId int) (err error) {
	err = utils.DB.Debug().Where("survey_id=?", surveyId).Delete(&entity.Question{}).Error
	optionService.DelBySurveyId(surveyId)
	return
}

func (ts *QuestionService) Get(id int) (q entity.Question, err error) {
	q.Id = id
	err = utils.DB.Find(&q).Error
	ops := optionService.List(id)
	q.Options = ops
	return
}
