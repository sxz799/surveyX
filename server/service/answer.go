package service

import (
	"errors"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
	"time"
)

type AnswerService struct {
}

func (as *AnswerService) List(a entity.AnswerSearch) (response.PageResult, error) {
	var answers []entity.Answer
	db := utils.DB.Model(&entity.Answer{})
	answer := a.Answer
	pi := a.PageInfo
	limit := pi.PageSize
	offset := pi.PageSize * (pi.PageNum - 1)
	var total int64
	if answer.SurveyId != "" {
		db = db.Where("survey_id = ?", answer.SurveyId)
	}
	if answer.QuestionId != 0 {
		db = db.Where("question_id = ?", answer.QuestionId)
	}
	if answer.Contact != "" {
		db = db.Where("contact = ?", answer.Contact)
	}
	if answer.Finger != "" {
		db = db.Where("finger = ?", answer.Finger)
	}
	db.Count(&total)
	err := db.Limit(limit).Offset(offset).Find(&answers).Error
	return response.PageResult{
		List:     answers,
		Total:    total,
		PageNum:  pi.PageNum,
		PageSize: pi.PageSize}, err
}

func (as *AnswerService) Add(ans []entity.Answer) (err error) {
	var survey entity.Survey
	var surveyId string
	var contact, finger string
	if len(ans) > 0 {
		surveyId = ans[0].SurveyId
		err = utils.DB.Where("id = ?", surveyId).First(&survey).Error
		if err != nil {
			return
		}
		contact = ans[0].Contact
		finger = ans[0].Finger
	}
	repeat := survey.Repeat
	check := survey.RepeatCheck
	switch repeat {
	case "no":
		switch check {
		case "contact":
			var num int64
			utils.DB.Debug().Model(&entity.Answer{}).Where("contact = ? and survey_id=?", contact, surveyId).Count(&num)
			if num > 0 {
				return errors.New("该问卷不允许重复提交")
			}
		case "finger":
			var num int64
			utils.DB.Model(&entity.Answer{}).Where("finger = ? and survey_id=?", finger, surveyId).Count(&num)
			if num > 0 {
				return errors.New("该问卷不允许重复提交")
			}
		}
	case "update":
		switch check {
		case "contact":
			utils.DB.Model(&entity.Answer{}).Delete(&entity.Answer{}, "contact = ? and survey_id=?", contact, surveyId)
		case "finger":
			utils.DB.Model(&entity.Answer{}).Delete(&entity.Answer{}, "finger = ? and survey_id=?", finger, surveyId)
		}
	}
	unix := time.Now().Unix()
	for _, a := range ans {
		a.Key = unix
		a.CreateAt = time.Now()
		err = utils.DB.Create(&a).Error
	}
	return
}
