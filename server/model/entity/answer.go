package entity

import (
	"github.com/sxz799/surveyX/model/common/request"
	"time"
)

type Answer struct {
	Id         int       `gorm:"primary_key" json:"id"`
	SurveyId   string    `json:"survey_id" form:"survey_id"`
	QuestionId int       `json:"question_id" form:"question_id"`
	Content    string    `json:"content" form:"content"`
	Label      string    `json:"label" form:"label"`
	ExtMsg     string    `json:"ext_msg" form:"ext_msg"`
	Contact    string    `json:"contact" form:"contact"`
	Finger     string    `json:"finger" form:"finger"`
	CreateAt   time.Time `json:"create_at" form:"create_at"`
	Key        string    `json:"key"`
}

type AnswerSearch struct {
	PageInfo request.PageInfo `json:"pageInfo"`
	Answer   Answer           `json:"answer"`
}
