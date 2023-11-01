package service

import (
	"github.com/sxz799/surveyX/model/common/request"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
)

type SurveyService struct {
}

var questionService QuestionService

func (ts *SurveyService) List(pi request.PageInfo) (response.PageResult, error) {
	var surveys []entity.Survey
	var total int64
	limit := pi.PageSize
	offset := pi.PageSize * (pi.PageNum - 1)
	db := utils.DB.Model(&entity.Survey{})
	db.Count(&total)
	db = db.Limit(limit).Offset(offset)
	err := db.Debug().Order("id DESC").Find(&surveys).Error
	return response.PageResult{
		List:     surveys,
		Total:    total,
		Page:     pi.PageNum,
		PageSize: pi.PageSize}, err
}

func (ts *SurveyService) Add(s entity.Survey) (err error) {
	err = utils.DB.Debug().Create(&s).Error
	return
}

func (ts *SurveyService) Update(s entity.Survey) error {
	err := utils.DB.Debug().Where("id=?", s.Id).Updates(&s).Error
	return err
}

func (ts *SurveyService) Del(id int) (err error) {
	s := entity.Survey{
		Id: id,
	}
	_ = questionService.DelBySurveyId(id)
	err = utils.DB.Delete(&s).Error
	return
}

func (ts *SurveyService) Get(id int) (s entity.Survey, err error) {
	s.Id = id
	err = utils.DB.First(&s).Error
	return
}
