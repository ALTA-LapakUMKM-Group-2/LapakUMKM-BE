package auth

import (
	"lapakUmkm/features/users"
)

type AuthServiceInterface interface {
	Login(email, password string) (string, users.UserEntity, error)
	Register(request users.UserEntity) error
	ChangePassword(id uint, old, new, confirm string) error

	GetSSOGoogleUrl() string
	LoginSSOGoogle(userEntity users.UserEntity) (string, users.UserEntity, error)
	ForgetPassword(email string) error
	IsUserExist(email string) error
	NewPassword(token, newPassword, confirmPassword string) error
}

type AuthDataInterface interface {
	GetUserByEmailOrId(email string, id uint) (users.UserEntity, error)
	Register(request users.UserEntity) error
	EditPassword(id uint, newPassword string) error
	EditData(request users.UserEntity) error
}
