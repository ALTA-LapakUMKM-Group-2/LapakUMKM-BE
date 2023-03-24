package service

import (
	"errors"
	"lapakUmkm/features/products"

	"github.com/go-playground/validator/v10"
)

type productService struct {
	Data     products.ProductDataInterface
	validate *validator.Validate
}

func New(data products.ProductDataInterface) products.ProductServiceInterface {
	return &productService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *productService) GetAll() ([]products.ProductEntity, error) {
	return s.Data.SelectAll()
}

func (s *productService) GetById(id uint) (products.ProductEntity, error) {
	return s.Data.SelectById(id)
}

func (s *productService) Create(productEntity products.ProductEntity) (products.ProductEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(productEntity, "User", "Category")
	if errValidate != nil {
		return products.ProductEntity{}, errValidate
	}

	productId, err := s.Data.Store(productEntity)
	if err != nil {
		return products.ProductEntity{}, err
	}

	return s.Data.SelectById(productId)
}

func (s *productService) Update(productEntity products.ProductEntity, id, userId uint) (products.ProductEntity, error) {
	checkDataExist, errData := s.Data.SelectById(id)
	if errData != nil {
		return checkDataExist, errData
	}

	if checkDataExist.UserId != userId {
		return products.ProductEntity{}, errors.New("can't update this product id")
	}

	err := s.Data.Edit(productEntity, id)
	if err != nil {
		return products.ProductEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *productService) Delete(id, userId uint) error {
	checkDataExist, err := s.Data.SelectById(id)
	if err != nil {
		return err
	}

	if checkDataExist.UserId != userId {
		return errors.New("don't have access to delete this product id")
	}

	return s.Data.Destroy(id)
}

// GetProductByCategoryId implements products.ProductServiceInterface
func (*productService) GetProductByCategoryId(userId uint) {
	panic("unimplemented")
}

// GetProductByUserId implements products.ProductServiceInterface
func (*productService) GetProductByUserId(userId uint) {
	panic("unimplemented")
}
