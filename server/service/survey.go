package service

import (
	"github.com/sxz799/surveyX/model/common/request"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
)

type SurveyService struct {
}

func (ts *SurveyService) List(pi request.PageInfo) (response.PageResult, error) {
	var surveys []entity.Survey
	var total int64
	limit := pi.PageSize
	offset := pi.PageSize * (pi.Page - 1)
	db := utils.DB.Model(&entity.Survey{})
	db.Count(&total)
	db = db.Limit(limit).Offset(offset)
	err := db.Order("id DESC").Find(&surveys).Error
	return response.PageResult{
		List:     surveys,
		Total:    total,
		Page:     pi.Page,
		PageSize: pi.PageSize}, err
}

func (ts *SurveyService) Add(s entity.Survey) (err error) {
	err = utils.DB.Debug().Create(&s).Error
	return
}

func (ts *SurveyService) Update(s entity.Survey) {
	utils.DB.Debug().Updates(&s)
}

func (ts *SurveyService) Del(id int) (err error) {
	s := entity.Survey{
		Id: id,
	}
	err = utils.DB.Delete(&s).Error
	return
}

func (ts *SurveyService) Get(id int) (s entity.Survey) {
	s.Id = id
	utils.DB.Find(&s)
	return
}
