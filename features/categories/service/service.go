package service

import (
	"lapakUmkm/features/categories"

	"github.com/go-playground/validator/v10"
)

type CategoryService struct {
	Data     categories.CategoryDataInterface
	validate *validator.Validate
}

func New(data categories.CategoryDataInterface) categories.CategoryServiceInterface {
	return &CategoryService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *CategoryService) GetAll() ([]categories.CategoryEntity, error) {
	return s.Data.SelectAll()
}

func (s *CategoryService) GetById(id uint) (categories.CategoryEntity, error) {
	return s.Data.SelectById(id)
}

func (s *CategoryService) Create(categoryEntity categories.CategoryEntity) (categories.CategoryEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(categoryEntity, "User")
	if errValidate != nil {
		return categories.CategoryEntity{}, errValidate
	}

	user_id, err := s.Data.Store(categoryEntity)
	if err != nil {
		return categories.CategoryEntity{}, err
	}

	return s.Data.SelectById(user_id)
}

func (s *CategoryService) Update(categoryEntity categories.CategoryEntity, id uint) (categories.CategoryEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(categoryEntity, "User")
	if errValidate != nil {
		return categories.CategoryEntity{}, errValidate
	}

	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	err := s.Data.Edit(categoryEntity, id)
	if err != nil {
		return categories.CategoryEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *CategoryService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}
