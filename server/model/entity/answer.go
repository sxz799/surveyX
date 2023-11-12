package entity

import (
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
}

type AnswerSearch struct {
	Answer    Answer    `json:"answer"`
	StartTime time.Time `json:"start_time" form:"start_time"`
	EndTime   time.Time `json:"end_time" form:"end_time"`
}
