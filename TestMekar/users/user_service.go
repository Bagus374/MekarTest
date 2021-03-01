package users

import "TestMekar/model"

type UserService interface {
	SignUp (user *model.User) (*model.User, error)
	FindAll() ([]*model.User, error)
	GetUserByID(id int) (*model.User,error)

}