package service

import (
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
)

type UserService struct{}

func (us *UserService) Add(user entity.User) (err error) {
	err = utils.DB.Create(&user).Error
	return
}

func (us *UserService) Update(user entity.User) (err error) {
	err = utils.DB.Save(&user).Error
	return
}

func (us *UserService) Delete(user entity.User) (err error) {
	err = utils.DB.Delete(&user).Error
	return
}

func (us *UserService) List(user entity.User) (users []entity.User, err error) {
	err = utils.DB.Find(&users).Error
	return
}

func (us *UserService) Get(user entity.User) (u entity.User, err error) {
	err = utils.DB.First(&u, user.Id).Error
	return
}

func (us *UserService) Login(user entity.User) (u entity.User, err error) {
	err = utils.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&u).Error
	return
}
