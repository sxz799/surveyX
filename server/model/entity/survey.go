package entity

import (
	"github.com/sxz799/surveyX/model/common/request"
	"time"
)

type Survey struct {
	Id          string    `gorm:"primary_key" json:"id"`
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Status      string    `json:"status" form:"status"`
	StartTime   time.Time `json:"start_time" form:"start_time"`
	EndTime     time.Time `json:"end_time" form:"end_time"`
	Questions   []int     `gorm:"-" json:"questions" form:"questions"`
	NeedContact string    `json:"need_contact" form:"need_contact"`
	ContactType string    `json:"contact_type" form:"contact_type"`
	Repeat      string    `json:"repeat" form:"repeat"`             // 是否允许重复提交 no yes update
	RepeatCheck string    `json:"repeat_check" form:"repeat_check"` // 重复提交检查方式 finger contact
	WaterMark   string    `json:"water_mark" form:"water_mark"`     // 水印
}

type SurveySearch struct {
	PageInfo request.PageInfo `json:"pageInfo"`
	Survey   Survey           `json:"survey"`
}
