package service

import (
	"errors"
	"lapakUmkm/features/users"
	"lapakUmkm/mocks"
	"lapakUmkm/utils/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var empty = users.UserEntity{}

func TestLogin(t *testing.T) {
	repo := mocks.NewAuthDataInterface(t)

	t.Run("loginNullInput", func(t *testing.T) {
		srv := New(repo)
		_, _, err := srv.Login("", "")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "email and password must be fill")
		repo.AssertExpectations(t)
	})

	hash, _ := helpers.HashPassword("123456")
	input := users.UserEntity{
		Id:       uint(1),
		Email:    "tes@gmail.com",
		Password: hash,
	}

	t.Run("error", func(t *testing.T) {
		repo.On("GetUserByEmailOrId", "tes2@gmail.com", uint(0)).Return(input, errors.New("user and password not found")).Once()
		srv := New(repo)
		_, _, err := srv.Login("tes2@gmail.com", "0812345")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "user and password not found")
		repo.AssertExpectations(t)
	})

	// t.Run("success", func(t *testing.T) {
	// 	repo.On("GetUserByEmailOrId", "tes@gmail.com", uint(0)).Return(input, nil).Once()
	// 	srv := New(repo)
	// 	token, userID, err := srv.Login("tes@gmail.com", "123456")
	// 	assert.Nil(t, err)
	// 	assert.NotEmpty(t, token)
	// 	assert.Equal(t, userID, input.Id)
	// 	repo.AssertExpectations(t)
	// })

}

func TestRegister(t *testing.T) {
	repo := mocks.NewAuthDataInterface(t)

	input := users.UserEntity{
		FullName: "Findryan",
		Email:    "tes@gmail.com",
		Password: "123456",
	}

	t.Run("success", func(t *testing.T) {
		srv := New(repo)
		input.Role = "user"
		repo.On("Register", input).Return(nil)
		err := srv.Register(input)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("errorValidation", func(t *testing.T) {
		srv := New(repo)
		err := srv.Register(empty)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "required")
		repo.AssertExpectations(t)
	})

	t.Run("errVal", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(errors.New("duplicated"))
		srv := New(repo)
		input.Id = uint(1)
		input.Role = "user"
		err := srv.Register(input)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "duplicated")
		repo.AssertExpectations(t)
	})
}

func TestChangePassword(t *testing.T) {
	repo := mocks.NewAuthDataInterface(t)
	// pass, _ := helpers.HashPassword("123456")
	expected := users.UserEntity{
		Id:       uint(1),
		FullName: "fin",
		Email:    "tes@gmail.com",
		Password: "123456",
	}

	t.Run("notNull", func(t *testing.T) {
		srv := New(repo)
		err := srv.ChangePassword(uint(1), "", "", "")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "cannot be empty")
		repo.AssertExpectations(t)
	})

	t.Run("notSamePassword", func(t *testing.T) {
		srv := New(repo)
		err := srv.ChangePassword(uint(1), "123456", "1234567", "1234568")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "must be similarity")
		repo.AssertExpectations(t)
	})

	t.Run("notExist", func(t *testing.T) {
		repo.On("GetUserByEmailOrId", ".", uint(1)).Return(expected, nil).Once()
		srv := New(repo)
		err := srv.ChangePassword(uint(1), "123456", "1234567", "1234567")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "old password not match with exist password")
		repo.AssertExpectations(t)
	})

	// t.Run("success", func(t *testing.T) {
	// 	repo.On("GetUserByEmailOrId", ".", uint(1)).Return(expected, nil).Once()
	// 	pass, _ := helpers.HashPassword("123456")
	// 	repo.On("EditPassword", uint(1), pass).Return(nil).Once()
	// 	srv := New(repo)
	// 	err := srv.ChangePassword(uint(1), "123456", "123456", "123456")
	// 	assert.NoError(t, err)
	// 	repo.AssertExpectations(t)
	// })
}

func TestIsUserExist(t *testing.T) {
	repo := mocks.NewAuthDataInterface(t)
	t.Run("notExist", func(t *testing.T) {
		repo.On("GetUserByEmailOrId", "haha@mail.com", uint(0)).Return(users.UserEntity{}, errors.New("email not found")).Once()
		srv := New(repo)
		err := srv.IsUserExist("haha@mail.com")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "email not found")
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("GetUserByEmailOrId", "haha@mail.com", uint(0)).Return(users.UserEntity{}, nil).Once()
		srv := New(repo)
		err := srv.IsUserExist("haha@mail.com")
		assert.Equal(t, err, nil)
		repo.AssertExpectations(t)
	})
}

func TestForgetPasswords(t *testing.T) {
	repo := mocks.NewAuthDataInterface(t)

	email := "haha@mail.com"
	t.Run("notExist", func(t *testing.T) {
		repo.On("GetUserByEmailOrId", "haha@mail.com", uint(0)).Return(users.UserEntity{}, errors.New("email not found")).Once()
		srv := New(repo)
		err := srv.IsUserExist("haha@mail.com")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "email not found")
		repo.AssertExpectations(t)
	})

	token := helpers.EncryptText(email)
	urlLink := "https://lapakumkm.netlify.app/new-password?token=" + token

	t.Run("errSendEmail", func(t *testing.T) {
		repo.On("GetUserByEmailOrId", email, uint(0)).Return(users.UserEntity{}, nil).Once()
		srv := New(repo)
		srv.ForgetPassword(email)
		err := helpers.SendMail("Forget Password", "", urlLink)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "title and email must be fill")
		repo.AssertExpectations(t)
	})
}

func TestNewPassword(t *testing.T) {
	repo := mocks.NewAuthDataInterface(t)
	t.Run("val", func(t *testing.T) {
		pass := "123456"
		srv := New(repo)
		err := srv.NewPassword(pass, pass, "")
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, "token,new password, and confirm password cannot be empty")
		repo.AssertExpectations(t)
	})
	pass := "123456"
	t.Run("notSame", func(t *testing.T) {
		srv := New(repo)
		err := srv.NewPassword(pass, pass, "123")
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, "new password and confirm password must be similarity")
		repo.AssertExpectations(t)
	})

	t.Run("notValid", func(t *testing.T) {
		srv := New(repo)
		email := helpers.DecryptText(pass)
		repo.On("GetUserByEmailOrId", email, uint(0)).Return(users.UserEntity{}, errors.New("not valid")).Once()
		err := srv.NewPassword(pass, pass, pass)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, "not valid")
		repo.AssertExpectations(t)
	})

}

func TestLoginSSOGoogle(t *testing.T) {
	// repo := mocks.NewAuthDataInterface(t)

	// request := users.UserEntity{
	// 	Email:        "tes@gmail.com",
	// 	PhotoProfile: "photo1",
	// 	FullName:     "tes@gmail.com",
	// 	Password:     "google-password",
	// 	Role:         "user",
	// }

	// // t.Run("error", func(t *testing.T) {
	// // 	repo.On("GetUserByEmailOrId", "tes2@gmail.com", uint(0)).Return(request, errors.New("user is not active anymore")).Once()
	// // 	srv := New(repo)
	// // 	_, _, err := srv.Login("tes@gmail.com", "google-password")
	// // 	assert.NotEmpty(t, err)
	// // 	assert.ErrorContains(t, err, "user and password not found")
	// // 	repo.AssertExpectations(t)
	// // })

	// // t.Run("notValid", func(t *testing.T) {
	// // 	srv := New(repo)
	// // 	_, _, err := srv.LoginSSOGoogle(request)
	// // 	assert.NotEmpty(t, err)
	// // 	assert.ErrorContains(t, err, "required")
	// // 	repo.AssertExpectations(t)
	// // })
}
