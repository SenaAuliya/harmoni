package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null; column:username"`
	Password       string `gorm:"column:password"`
	PhoneNumber    string
	ProfilePicture string
	Bio            string
	Status         string
}
