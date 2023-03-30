package data

import (
	"lapakUmkm/utils/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName     string
	Email        string `gorm:"unique"`
	Password     string
	Role         string
	PhoneNumber  string
	Address      string
	ShopName     string
	PhotoProfile string
	Status       string
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password, err = helpers.HashPassword(user.Password)
	user.Status = "inactive"
	return
}
