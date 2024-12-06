package common

import (
	"github.com/sxz799/surveyX/model/entity"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}
func (s *Service) Login(user entity.User) (u entity.User, err error) {
	err = s.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&u).Error
	return
}
