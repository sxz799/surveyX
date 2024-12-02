package user

import (
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
)

type Service struct{}

func (us *Service) Add(user entity.User) (id int, err error) {
	err = utils.DB.Create(&user).Error
	id = user.Id
	return
}

func (us *Service) Update(user entity.User) (err error) {
	err = utils.DB.Updates(&user).Error
	return
}

func (us *Service) Delete(user entity.User) (err error) {
	err = utils.DB.Delete(&user).Error
	return
}

func (us *Service) List(user entity.User) (users []entity.User, err error) {
	err = utils.DB.Model(&user).Find(&users).Error
	return
}

func (us *Service) GetByGithubId(id int) (u entity.User, err error) {
	err = utils.DB.Where("github_uid=?", id).First(&u).Error
	return
}

func (us *Service) Login(user entity.User) (u entity.User, err error) {
	err = utils.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&u).Error
	return
}
