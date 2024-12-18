package service

import (
	"github.com/sxz799/surveyX/service/answer"
	"github.com/sxz799/surveyX/service/common"
	"github.com/sxz799/surveyX/service/question"
	"github.com/sxz799/surveyX/service/survery"
	"github.com/sxz799/surveyX/service/user"
	"gorm.io/gorm"
)

type Container struct {
	AnswerService   *answer.Service
	OptionService   *question.OptionService
	QuestionService *question.Service
	SurveyService   *survery.Service
	UserService     *user.Service
	CommonService   *common.Service
}

func NewContainer(db *gorm.DB) *Container {
	answerService := answer.NewService(db)
	optionService := question.NewOptionService(db)
	questionService := question.NewService(db, optionService, answerService)
	surveyService := survery.NewService(db, questionService)
	userService := user.NewService(db)
	commonService := common.NewService(db)
	return &Container{
		AnswerService:   answerService,
		OptionService:   optionService,
		QuestionService: questionService,
		SurveyService:   surveyService,
		UserService:     userService,
		CommonService:   commonService,
	}
}
