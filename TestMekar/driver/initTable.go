package driver

import (
	"TestMekar/model"
	"github.com/jinzhu/gorm"
)

func InitTable(db *gorm.DB)  {
	db.Debug().AutoMigrate(
		&model.User{})
}
