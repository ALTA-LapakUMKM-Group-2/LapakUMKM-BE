package service

import (
	"errors"
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/auth"
	"lapakUmkm/features/users"
	"lapakUmkm/utils/helpers"

	"github.com/go-playground/validator/v10"
)

type authService struct {
	data     auth.AuthDataInterface
	validate *validator.Validate
}

func New(u auth.AuthDataInterface) auth.AuthServiceInterface {
	return &authService{
		data:     u,
		validate: validator.New(),
	}
}

func (u *authService) Login(email, password string) (string, users.UserEntity, error) {
	if email == "" || password == "" {
		return "", users.UserEntity{}, errors.New("email and password must be fill")
	}

	user, err := u.data.GetUserByEmailOrId(email, 0)
	if err != nil || !helpers.CheckPasswordHash(password, user.Password) {
		return "", users.UserEntity{}, errors.New("user and password not found")
	}

	//make jwt
	token, errToken := middlewares.CreateToken(int(user.Id), user.Role)
	if errToken != nil {
		return "", users.UserEntity{}, errToken
	}

	return token, user, nil
}

func (h *authService) Register(request users.UserEntity) error {
	h.validate = validator.New()
	if errValidate := h.validate.Struct(request); errValidate != nil {
		return errValidate
	}
	return h.data.Register(request)
}

func (u *authService) ChangePassword(id uint, oldPassword, newPassword, confirmPssword string) error {
	if oldPassword == "" || newPassword == "" || confirmPssword == "" {
		return errors.New("old password,new password, and confirm password cannot be empty")
	}

	if newPassword != confirmPssword {
		return errors.New("new password and confirm password must be similarity")
	}

	user, err := u.data.GetUserByEmailOrId(".", id)
	if err != nil || !helpers.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("old password not match with exist password")
	}

	hash, _ := helpers.HashPassword(newPassword)
	return u.data.EditPassword(id, hash)
}
