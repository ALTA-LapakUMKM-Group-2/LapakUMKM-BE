package auth

import (
	"lapakUmkm/features/users"
)

type AuthServiceInterface interface {
	Login(email, password string) (string, users.UserEntity, error)
	Register(request users.UserEntity) error
	ChangePassword(id uint, old, new, confirm string) error
}

type AuthDataInterface interface {
	GetUserByEmailOrId(email string, id uint) (users.UserEntity, error)
	Register(request users.UserEntity) error
	EditPassword(id uint, newPassword string) error
}
