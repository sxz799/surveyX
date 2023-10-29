package service

import (
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
)

type OptionService struct {
}

func (ts *OptionService) List(qid int) (ops []entity.Option) {
	db := utils.DB.Model(&entity.Option{})
	db.Where("question_id=?", qid).Find(&ops)
	return
}

func (ts *OptionService) Add(ops []entity.Option) {
	for _, op := range ops {
		utils.DB.Create(&op)
	}
}

func (ts *OptionService) Del(questionId int) {

	utils.DB.Where("question_id=?", questionId).Delete(&entity.Option{})

}
