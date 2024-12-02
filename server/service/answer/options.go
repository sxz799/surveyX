package answer

import (
	"github.com/sxz799/surveyX/model/entity"
	"gorm.io/gorm"
)

type OptionService struct {
	db *gorm.DB
}

func NewOptionService(db *gorm.DB) *OptionService {
	return &OptionService{db: db}
}

func (s *OptionService) List(qid int) (ops []entity.Option) {
	db := s.db.Model(&entity.Option{})
	db.Where("question_id=?", qid).Find(&ops)
	return
}

func (s *OptionService) Add(ops []entity.Option) {
	for _, op := range ops {
		s.db.Create(&op)
	}
}

func (s *OptionService) Del(questionId int) {

	s.db.Where("question_id=?", questionId).Delete(&entity.Option{})

}

func (s *OptionService) DelBySurveyId(sId string) {

	s.db.Where("survey_id=?", sId).Delete(&entity.Option{})

}
