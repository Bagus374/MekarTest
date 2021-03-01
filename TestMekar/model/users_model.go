package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"column:name;size:255;not null" json:"name"`
	DateOfBirth string `gorm:"column:date;type:date;not null" json:"date_of_birth"`
	NoKtp       string `gorm:"column:noKtp; not null" json:"no_ktp"`
	Pekerjaan   string `gorm:"column:pekerjaan; not null" json:"pekerjaan"`
	Pendidikan  string `gorm:"column:pendidikan; not null" json:"pendidikan"`
}
