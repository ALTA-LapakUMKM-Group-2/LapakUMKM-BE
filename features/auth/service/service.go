package service

import (
	"errors"
	"fmt"
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

func (s *authService) GetSSOGoogleUrl() string {
	return helpers.GetSSOGoogleUrl()
}

func (s *authService) LoginSSOGoogle(userEntity users.UserEntity) (string, users.UserEntity, error) {
	request := users.UserEntity{
		Email:        userEntity.Email,
		PhotoProfile: userEntity.PhotoProfile,
		FullName:     userEntity.Email,
		Password:     "google-password",
		Role:         "user",
	}

	user, _ := s.data.GetUserByEmailOrId(userEntity.Email, 0)
	s.validate = validator.New()
	errValidate := s.validate.Struct(request)
	if errValidate != nil {
		return "", users.UserEntity{}, errValidate
	}

	if user.Id == 0 {
		s.data.Register(request)
	} else {
		request.Id = user.Id
		request.Password = user.Password
		request.Role = user.Role
		request.FullName = user.FullName
	}
	//edit data from google
	if err := s.data.EditData(userEntity); err != nil {
		return "", users.UserEntity{}, err
	}

	//login
	userLogin, _ := s.data.GetUserByEmailOrId(userEntity.Email, 0)
	token, errToken := middlewares.CreateToken(int(userLogin.Id), userLogin.Role)
	if errToken != nil {
		return "", users.UserEntity{}, errToken
	}

	return token, userLogin, nil
}

// cekemailexist
func (s *authService) IsUserExist(email string) error {
	if _, err := s.data.GetUserByEmailOrId(email, 0); err != nil {
		return errors.New("email not found")
	}
	return nil
}

func (s *authService) ForgetPassword(email string) error {
	if _, err := s.data.GetUserByEmailOrId(email, 0); err != nil {
		return errors.New("email not found")
	}
	token := helpers.EncryptText(email)
	urlLink := "https://lapak-umkm-test-pase1.vercel.app/new-password?token=" + token

	//send URL to Email
	if errSendmail := helpers.SendMail(email, urlLink); errSendmail != nil {
		return errSendmail
	}

	return nil
}

func (s *authService) NewPassword(token, newPassword, confirmPssword string) error {
	if token == "" || newPassword == "" || confirmPssword == "" {
		return errors.New("token,new password, and confirm password cannot be empty")
	}

	if newPassword != confirmPssword {
		return errors.New("new password and confirm password must be similarity")
	}

	email := helpers.DecryptText(token)
	fmt.Println(email)
	user, err := s.data.GetUserByEmailOrId(email, 0)
	if err != nil {
		return errors.New("not valid")
	}
	hash, _ := helpers.HashPassword(newPassword)
	return s.data.EditPassword(user.Id, hash)
}
