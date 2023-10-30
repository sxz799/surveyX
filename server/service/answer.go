package service

import (
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
)

type AnswerService struct {
}

func (ts *AnswerService) Add(as []entity.Answer) (err error) {
	for _, a := range as {
		err = utils.DB.Create(&a).Error
	}
	return
}
