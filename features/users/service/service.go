package service

import (
	"errors"
	"lapakUmkm/features/users"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	Data     users.UserDataInterface
	validate *validator.Validate
}

func New(data users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		Data: data,
	}
}

func (s *userService) GetUser(id uint) (users.UserEntity, error) {
	return s.Data.SelectById(id)
}

func (s *userService) GetAll() ([]users.UserEntity, error) {
	return s.Data.SelectAll()
}

func (s *userService) GetById(id uint) (users.UserEntity, error) {
	return s.Data.SelectById(id)
}

func (s *userService) Create(userEntity users.UserEntity) (users.UserEntity, error) {
	if userEntity.Role != "admin" && userEntity.Role != "user" {
		return users.UserEntity{}, errors.New("role option only : admin and user")
	}

	s.validate = validator.New()
	errValidate := s.validate.StructExcept(userEntity, "Team")
	if errValidate != nil {
		return users.UserEntity{}, errValidate
	}

	user_id, err := s.Data.Store(userEntity)
	if err != nil {
		return users.UserEntity{}, err
	}

	return s.Data.SelectById(user_id)
}

func (s *userService) Update(id uint, request users.UserEntity) (users.UserEntity, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	_, err := s.Data.Edit(request, id)
	if err != nil {
		return users.UserEntity{}, err
	}

	return s.Data.SelectById(id)
}

func (s *userService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}

func (s *userService) UpdateToSeller(id uint, request users.UserEntity) (users.UserEntity, error) {
	//cek all data
	usersData, _ := s.Data.SelectById(id)
	if usersData.Address == "" || usersData.PhoneNumber == "" {
		return users.UserEntity{}, errors.New("complete all your data first. (address and phone number)")
	}

	if request.ShopName == "" {
		return users.UserEntity{}, errors.New("insert shop name")
	}

	//update to seller
	request.Role = "seller"
	if _, err := s.Data.Edit(request, id); err != nil {
		return users.UserEntity{}, err
	}

	return s.Data.SelectById(id)
}

func (s *userService) UpdateToProfile(id uint, request users.UserEntity) (users.UserEntity, error) {
	panic("unimplemented")
}
