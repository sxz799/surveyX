package entity

import "github.com/sxz799/surveyX/model/common/request"

type Question struct {
	Id       int      `gorm:"primary_key" json:"id"`
	SurveyId string   `json:"survey_id" form:"survey_id"`
	Text     string   `json:"text" form:"text"`
	Type     string   `json:"type" form:"type"`
	Options  []Option `gorm:"-" json:"options" form:"options"`
	Order    int      `json:"order" form:"order"`
}

type Option struct {
	Id         int    `gorm:"primary_key" json:"id"`
	QuestionId int    `json:"question_id" form:"question_id"`
	SurveyId   string `json:"survey_id" form:"survey_id"`
	Label      string `json:"label" form:"label"`
	Value      string `json:"value" form:"value"`
	HasExtMsg  string `json:"has_ext_msg" form:"has_ext_msg"`
}

type QuestionSearch struct {
	PageInfo request.PageInfo `json:"pageInfo"`
	Question Question         `json:"question"`
}
