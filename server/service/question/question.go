package question

import (
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/service/answer"
	"github.com/sxz799/surveyX/utils"
)

type Service struct{}

var optionsService answer.OptionService
var answerService answer.Service

func (ts *Service) List(q entity.QuestionSearch) (response.PageResult, error) {
	var qs []entity.Question
	var total int64
	pi := q.PageInfo
	sId := q.Question.SurveyId
	limit := pi.PageSize
	offset := pi.PageSize * (pi.PageNum - 1)
	db := utils.DB.Model(&entity.Question{})
	db.Where("survey_id=?", sId)
	db.Count(&total)
	err := db.Limit(limit).Offset(offset).Order("`order`").Find(&qs).Error
	for i := range qs {
		ops := optionsService.List(qs[i].Id)
		qs[i].Options = ops
	}
	return response.PageResult{
		List:     qs,
		Total:    total,
		PageNum:  pi.PageNum,
		PageSize: pi.PageSize}, err
}

func (ts *Service) Add(q entity.Question) (err error) {

	err = utils.DB.Create(&q).Error

	for i := range q.Options {
		q.Options[i].QuestionId = q.Id
		q.Options[i].SurveyId = q.SurveyId
	}

	optionsService.Add(q.Options)
	return
}

func (ts *Service) Update(q entity.Question) (err error) {

	err = utils.DB.Updates(&q).Error
	for i := range q.Options {
		q.Options[i].QuestionId = q.Id
		q.Options[i].SurveyId = q.SurveyId
	}
	optionsService.Del(q.Id)
	optionsService.Add(q.Options)
	return
}

func (ts *Service) Del(id int) (err error) {
	q := entity.Question{
		Id: id,
	}
	err = utils.DB.Delete(&q).Error
	optionsService.Del(id)
	answerService.DelByQuestionId(id)
	return
}

func (ts *Service) DelBySurveyId(surveyId string) (err error) {
	err = utils.DB.Delete(&entity.Question{}, "survey_id=?", surveyId).Error
	optionsService.DelBySurveyId(surveyId)
	return
}

func (ts *Service) Get(id int) (q entity.Question, err error) {
	q.Id = id
	err = utils.DB.Find(&q).Error
	ops := optionsService.List(id)
	q.Options = ops
	return
}

func (ts *Service) Analysis(id string) (any, error) {
	type TAnswer struct {
		QuestionId int    `json:"question_id" form:"question_id"`
		Content    string `json:"content" form:"content"`
		Label      string `json:"label" form:"label"`
		Contact    string `json:"contact" form:"contact"`
		Finger     string `json:"finger" form:"finger"`
		CreateAt   string `json:"create_at" form:"create_at"`
	}
	var answers []TAnswer
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
	var LabelCount []struct {
		Label string `json:"label"`
		Count int    `json:"count"`
	}
	for k, v := range labelMap {
		LabelCount = append(LabelCount, struct {
			Label string `json:"label"`
			Count int    `json:"count"`
		}{Label: k, Count: v})
	}
	r.ContactCount = len(contactMap)
	r.FingerCount = len(fingerMap)
	r.LabelInfo = LabelCount
	r.Count = len(answers)
	return r, err
}
