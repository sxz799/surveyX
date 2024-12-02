package user

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

func (s *Service) Add(user entity.User) (id int, err error) {
	err = s.db.Create(&user).Error
	id = user.Id
	return
}

func (s *Service) Update(user entity.User) (err error) {
	err = s.db.Updates(&user).Error
	return
}

func (s *Service) Delete(user entity.User) (err error) {
	err = s.db.Delete(&user).Error
	return
}

func (s *Service) List(user entity.User) (users []entity.User, err error) {
	err = s.db.Model(&user).Find(&users).Error
	return
}

func (s *Service) GetByGithubId(id int) (u entity.User, err error) {
	err = s.db.Where("github_uid=?", id).First(&u).Error
	return
}

func (s *Service) Login(user entity.User) (u entity.User, err error) {
	err = s.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&u).Error
	return
}
