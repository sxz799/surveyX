package entity

type Survey struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Questions   []int  `gorm:"-"`
}

type Question struct {
	Id       int      `gorm:"primary_key"`
	SurveyID int      `json:"surveyID"`
	Text     string   `json:"text"`
	Type     int      `json:"type"`
	Options  []Option `gorm:"-"`
	Order    int      `json:"order"`
}

type Option struct {
	Id         int    `gorm:"primary_key"`
	QuestionId int    `json:"questionId"`
	Label      string `json:"label"`
	Value      string `json:"value"`
}

type Answer struct {
	Id         int    `gorm:"primary_key"`
	SurveyId   int    `json:"surveyId"`
	QuestionId int    `json:"questionId"`
	Content    string `json:"content"`
}
