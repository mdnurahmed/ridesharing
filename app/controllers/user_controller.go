package controllers

import (
	"shopup/app/models"
	"shopup/app/repositories"
)

type IUserController interface {
	AddUser(models.User) error
	AddVechile(models.Vechile) error
}

type UserController struct {
	userRepo repositories.IUserRepo
}

func NewInstanceOfUserController(userRepo repositories.IUserRepo) UserController {
	return UserController{userRepo: userRepo}
}

func (u *UserController) AddUser(user models.User) error {
	return u.userRepo.AddUser(user)
}

func (u *UserController) AddVechile(vechile models.Vechile) error {
	return u.userRepo.AddVechile(vechile)
}
