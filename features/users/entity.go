package users

import (
	"time"
)

type UserEntity struct {
	Id           uint
	FullName     string `validate:"required"`
	Email        string `validate:"required,email"`
	Password     string `validate:"required"`
	PhoneNumber  string
	Address      string
	Role         string `validate:"required"`
	ShopName     string
	PhotoProfile string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserServiceInterface interface {
	GetUser(id uint) (UserEntity, error)
	Update(id uint, userEntity UserEntity) (UserEntity, error)
	Delete(id uint) error
	UpdateToProfile(id uint, userEntity UserEntity) (UserEntity, error)
	UpdateToSeller(id uint, userEntity UserEntity) (UserEntity, error)
}

type UserDataInterface interface {
	SelectAll() ([]UserEntity, error)
	SelectById(id uint) (UserEntity, error)
	Store(userEntity UserEntity) (uint, error)
	Edit(userEntity UserEntity, id uint) (uint, error)
	Destroy(id uint) error
}
