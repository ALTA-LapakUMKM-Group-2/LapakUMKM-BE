package service

import (
	"errors"
	"lapakUmkm/features/auth"
	"lapakUmkm/features/users"
	"lapakUmkm/mocks"
	"lapakUmkm/utils/helpers"
	"testing"

	"github.com/go-playground/validator/v10"
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
		FullName: "yayan",
	}

	t.Run("error", func(t *testing.T) {
		repo.On("GetUserByEmailOrId", "tes2@gmail.com", uint(0)).Return(input, errors.New("user and password not found")).Once()
		srv := New(repo)
		_, _, err := srv.Login("tes2@gmail.com", "0812345")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "user and password not found")
		repo.AssertExpectations(t)
	})
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
	expected := users.UserEntity{
		Id:       uint(1),
		FullName: "fin",
		Email:    "tes@gmail.com",
		Password: "123456",
	}

	// t.Run("success", func(t *testing.T) {
	// 	srv := New(repo)
	// 	password := "password123"
	// 	hashedPassword, _ := helpers.HashPassword(password)

	// 	repo.On("GetUserByEmailOrId", ".", uint(1)).Return(&users.UserEntity{
	// 		Id:       1,
	// 		Email:    "john@example.com",
	// 		Password: hashedPassword,
	// 	}, nil)

	// 	repo.On("EditPassword", uint(1), hashedPassword).Return(nil)

	// 	err := srv.ChangePassword(1, password, password, password)
	// 	assert.NoError(t, err)

	// 	repo.AssertExpectations(t)
	// })

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

func TestForgetPassword(t *testing.T) {
	email := "haha@mail.com"
	token := helpers.EncryptText(email)
	urlLink := "https://lapakumkm.netlify.app/new-password?token=" + token

	repo := mocks.NewAuthDataInterface(t)

	// t.Run("success", func(t *testing.T) {
	// 	user := users.UserEntity{
	// 		Id:       uint(123),
	// 		Email:    email,
	// 		Password: "hashedpassword",
	// 		// add other relevant fields as needed
	// 	}
	// 	repo.On("GetUserByEmailOrId", email, uint(0)).Return(user, nil).Once()
	// 	srv := New(repo)
	// 	err := srv.ForgetPassword(email)
	// 	assert.NoError(t, err)

	// 	errSendMail := helpers.SendMail("Forget Password", email, urlLink)
	// 	assert.NoError(t, errSendMail)

	// 	repo.AssertExpectations(t)
	// })

	t.Run("notExist", func(t *testing.T) {
		repo.On("GetUserByEmailOrId", "haha@mail.com", uint(0)).Return(users.UserEntity{}, errors.New("email not found")).Once()
		srv := New(repo)
		err := srv.ForgetPassword("haha@mail.com")
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "email not found")
		repo.AssertExpectations(t)
	})

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

	mockData := users.UserEntity{
		Id:       uint(1),
		Email:    helpers.EncryptText("test@example.com"),
		Password: "",
	}
	repo := mocks.NewAuthDataInterface(t)

	t.Run("success", func(t *testing.T) {
		pass := "newpassword"
		hash, _ := helpers.HashPassword(pass)
		mockData.Password = hash
		srv := New(repo)
		email := helpers.DecryptText(mockData.Email)
		repo.On("GetUserByEmailOrId", email, uint(0)).Return(mockData, nil).Once()
		repo.On("EditPassword", uint(1), mock.AnythingOfType("string")).Return(nil).Once()
		srv.NewPassword(mockData.Email, pass, pass)
		repo.AssertExpectations(t)

		err := helpers.CheckPasswordHash(pass, hash)
		assert.Equal(t, err, true)
	})
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

func Test_authService_ChangePassword(t *testing.T) {
	type fields struct {
		data     auth.AuthDataInterface
		validate *validator.Validate
	}
	type args struct {
		id             uint
		oldPassword    string
		newPassword    string
		confirmPssword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authService{
				data:     tt.fields.data,
				validate: tt.fields.validate,
			}
			if err := u.ChangePassword(tt.args.id, tt.args.oldPassword, tt.args.newPassword, tt.args.confirmPssword); (err != nil) != tt.wantErr {
				t.Errorf("authService.ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
