package service

import (
	"errors"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
	"log"
)

type AnswerService struct {
}

func (ts *AnswerService) Add(as []entity.Answer) (err error) {
	var survey entity.Survey
	var surveyId int
	if len(as) > 0 {
		surveyId = as[0].SurveyId
		err = utils.DB.Where("id = ?", surveyId).First(&survey).Error
		if err != nil {
			return
		}
	}
	repeat := survey.Repeat
	check := survey.RepeatCheck
	switch repeat {

	case "no":
		switch check {
		case "contact":
			contact := as[0].Contact
			var num int64
			utils.DB.Debug().Model(&entity.Answer{}).Where("contact = ? and survey_id=?", contact, surveyId).Count(&num)
			log.Println(num)
			if num > 0 {

				return errors.New("该问卷不允许重复提交")
			}
		case "finger":
			finger := as[0].Finger
			var num int64
			utils.DB.Model(&entity.Answer{}).Where("finger = ? and survey_id=?", finger, surveyId).Count(&num)
			if num > 0 {
				return errors.New("该问卷不允许重复提交")
			}
		}
	case "yes_but_update":
		switch check {
		case "contact":
			contact := as[0].Contact
			utils.DB.Model(&entity.Answer{}).Delete(&entity.Answer{}, "contact = ? and survey_id=?", contact, surveyId)
		case "finger":
			finger := as[0].Finger
			utils.DB.Model(&entity.Answer{}).Delete(&entity.Answer{}, "finger = ? and survey_id=?", finger, surveyId)
		}
	}
	for _, a := range as {
		err = utils.DB.Create(&a).Error
	}
	return
}
