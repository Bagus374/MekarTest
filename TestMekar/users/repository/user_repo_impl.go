package repository

import (
	"TestMekar/model"
	"TestMekar/users"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func (u UserRepoImpl) GetUserByID(id int) (*model.User, error) {
	dataUser := new(model.User)

	if err := u.db.Table("users").Where("id = ?", id).First(&dataUser).Error; err != nil {
		return nil, errors.New("ERROR: Error no data user with id you entered")
	}

	return dataUser, nil
}

func (u UserRepoImpl) FindAll() ([]*model.User, error) {
	userlist := make([]*model.User, 0)

	err := u.db.Table("users").Find(&userlist).Error
	if err != nil {
		fmt.Errorf("[UserRepoImpl.FindAll] Error when trying to get all data. error is : %v\n", err)
	}
	return userlist, nil
}

func (u UserRepoImpl) SignUp(user *model.User) (*model.User, error) {
	err := u.db.Table("users").Create(&user).Error
	if err != nil {
		fmt.Errorf("[UserRepoImpl.SignUp] Error when trying to create use error is : %v\n", err)
		return nil, err
	}
	return user, nil
}

func CreateRepoImpl(db *gorm.DB) users.UserRepository {
	return &UserRepoImpl{db}
}
