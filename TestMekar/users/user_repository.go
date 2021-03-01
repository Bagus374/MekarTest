package users

import (
	"TestMekar/model"
)

type UserRepository interface {
	SignUp (user *model.User) (*model.User, error)
	FindAll() ([]*model.User, error)
	GetUserByID(id int) (*model.User,error)
}
