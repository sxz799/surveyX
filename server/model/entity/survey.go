package entity

type Survey struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Questions   []int  `gorm:"-" json:"questions"`
}

type Question struct {
	Id       int      `gorm:"primary_key" json:"id"`
	SurveyId int      `json:"survey_id"`
	Text     string   `json:"text"`
	Type     string   `json:"type"`
	Options  []Option `gorm:"-" json:"options"`
	Order    int      `json:"order"`
}

type Option struct {
	Id         int    `gorm:"primary_key" json:"id"`
	QuestionId int    `json:"question_id"`
	Label      string `json:"label"`
	Value      string `json:"value"`
}

type Answer struct {
	Id         int    `gorm:"primary_key" json:"id"`
	SurveyId   int    `json:"survey_id"`
	QuestionId int    `json:"question_id"`
	Content    string `json:"content"`
}
