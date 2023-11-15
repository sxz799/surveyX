package service

import (
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
)

type QuestionService struct {
}

var optionService OptionService

func (ts *QuestionService) List(q entity.QuestionSearch) (response.PageResult, error) {
	var qs []entity.Question
	var total int64
	pi := q.PageInfo
	sId := q.Question.SurveyId
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
		PageNum:  pi.PageNum,
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

func (ts *QuestionService) DelBySurveyId(surveyId string) (err error) {
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

func (ts *QuestionService) Analysis(id string) (any, error) {
	type answer struct {
		QuestionId int    `json:"question_id" form:"question_id"`
		Content    string `json:"content" form:"content"`
		Label      string `json:"label" form:"label"`
		Contact    string `json:"contact" form:"contact"`
		Finger     string `json:"finger" form:"finger"`
		CreateAt   string `json:"create_at" form:"create_at"`
	}
	var answers []answer
	err := utils.DB.Table("answers").Where("question_id=?", id).Order("key").Find(&answers).Error
	contactMap := make(map[string]struct{})
	fingerMap := make(map[string]struct{})
	labelMap := make(map[string]int)
	for _, a := range answers {
		contactMap[a.Contact] = struct{}{}
		fingerMap[a.Finger] = struct{}{}
		if a.Label != "" {
			labelMap[a.Label]++
		}

	}
	type result struct {
		ContactCount int `json:"contact_count"`
		FingerCount  int `json:"finger_count"`
		LabelInfo    any `json:"label_info"`
		Count        int `json:"count"`
	}
	var r result
	r.ContactCount = len(contactMap)
	r.FingerCount = len(fingerMap)
	r.LabelInfo = labelMap
	r.Count = len(answers)
	return r, err
}
