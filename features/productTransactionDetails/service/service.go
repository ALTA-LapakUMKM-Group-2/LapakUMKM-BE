package service

import (
	"lapakUmkm/features/productTransactionDetails"

	"github.com/go-playground/validator/v10"
)

type ProductTransactionDetailsService struct {
	Data     productTransactionDetails.ProductTransactionDetailDataInterface
	validate *validator.Validate
}

func New(data productTransactionDetails.ProductTransactionDetailDataInterface) productTransactionDetails.ProductTransactionDetailServiceInterface {
	return &ProductTransactionDetailsService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *ProductTransactionDetailsService) Create(productTransactionDetailEntity productTransactionDetails.ProductTransactionDetailEntity) (productTransactionDetails.ProductTransactionDetailEntity, error) {
	id, err := s.Data.Store(productTransactionDetailEntity)
	if err != nil {
		return productTransactionDetails.ProductTransactionDetailEntity{}, err
	}

	return s.Data.SelectById(id)
}

func (s *ProductTransactionDetailsService) GetById(id uint) (productTransactionDetails.ProductTransactionDetailEntity, error) {
	return s.Data.SelectById(id)
}

func (s *ProductTransactionDetailsService) GetByTransaksiId(productId uint) ([]productTransactionDetails.ProductTransactionDetailEntity, error) {
	return s.Data.SelectByTransaksiId(productId)
}
