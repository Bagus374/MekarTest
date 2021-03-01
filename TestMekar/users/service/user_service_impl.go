package service

import (
	"TestMekar/model"
	"TestMekar/users"
)

type UserServiceImpl struct {
	userRepository users.UserRepository
}

func (u UserServiceImpl) GetUserByID(id int) (*model.User, error) {
	return u.userRepository.GetUserByID(id)
}

func (u UserServiceImpl) FindAll() ([]*model.User, error) {
	return u.userRepository.FindAll()
}

func (u UserServiceImpl) SignUp(user *model.User) (*model.User, error) {
	return u.userRepository.SignUp(user)
}

func CreateServiceImpl(userRepository users.UserRepository) users.UserService {
	return &UserServiceImpl{userRepository}
}
