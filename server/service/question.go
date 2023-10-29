package service

import (
	"github.com/sxz799/surveyX/model/common/request"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
	"log"
)

type QuestionService struct {
}

var optionService OptionService

func (ts *QuestionService) List(pi request.PageInfo, sId int) (response.PageResult, error) {
	var qs []entity.Question
	var total int64
	limit := pi.PageSize
	offset := pi.PageSize * (pi.Page - 1)
	db := utils.DB.Model(&entity.Question{})
	db.Count(&total)
	db = db.Limit(limit).Offset(offset)
	err := db.Where("survey_id=?", sId).Order("id DESC").Find(&qs).Error

	for _, q := range qs {
		ops := optionService.List(q.Id)
		q.Options = ops
	}

	return response.PageResult{
		List:     qs,
		Total:    total,
		Page:     pi.Page,
		PageSize: pi.PageSize}, err
}

func (ts *QuestionService) Add(q entity.Question) (err error) {

	err = utils.DB.Debug().Create(&q).Error

	log.Println("q.id:", q.Id)

	for i := range q.Options {
		q.Options[i].QuestionId = q.Id
	}

	optionService.Add(q.Options)
	return
}

func (ts *QuestionService) Update(q entity.Question) (err error) {

	err = utils.DB.Debug().Updates(&q).Error
	for i := range q.Options {
		q.Options[i].QuestionId = q.Id
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

func (ts *QuestionService) Get(id int) (q entity.Question, err error) {
	q.Id = id
	err = utils.DB.Find(&q).Error
	ops := optionService.List(id)
	q.Options = ops
	return
}
