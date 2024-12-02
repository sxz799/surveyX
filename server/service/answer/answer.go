package answer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) List(a entity.AnswerSearch) (response.PageResult, error) {
	var answers []entity.Answer
	db := s.db.Model(&entity.Answer{})
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

func (s *Service) Add(ans []entity.Answer) (err error) {
	var survey entity.Survey
	var surveyId string
	var contact, finger string
	if len(ans) > 0 {
		surveyId = ans[0].SurveyId
		err = s.db.Where("id = ?", surveyId).First(&survey).Error
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
			s.db.Model(&entity.Answer{}).Where("contact = ? and survey_id=?", contact, surveyId).Count(&num)
			if num > 0 {
				return errors.New("该问卷不允许重复提交")
			}
		case "finger":
			var num int64
			s.db.Model(&entity.Answer{}).Where("finger = ? and survey_id=?", finger, surveyId).Count(&num)
			if num > 0 {
				return errors.New("该问卷不允许重复提交")
			}
		}
	case "update":
		switch check {
		case "contact":
			s.db.Model(&entity.Answer{}).Delete(&entity.Answer{}, "contact = ? and survey_id=?", contact, surveyId)
		case "finger":
			s.db.Model(&entity.Answer{}).Delete(&entity.Answer{}, "finger = ? and survey_id=?", finger, surveyId)
		}
	}

	key := strings.ReplaceAll(uuid.New().String(), "-", "")

	for _, a := range ans {
		a.Key = key
		a.CreateAt = time.Now()
		err = s.db.Create(&a).Error
	}
	return
}

func (s *Service) DelBySurveyId(surveyId string) (err error) {
	err = s.db.Delete(&entity.Answer{}, "survey_id = ?", surveyId).Error
	return
}

func (s *Service) DelByQuestionId(questionId int) (err error) {
	err = s.db.Delete(&entity.Answer{}, "question_id = ?", questionId).Error
	return
}
