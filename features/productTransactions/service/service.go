package service

import (
	"lapakUmkm/features/productTransactions"

	"github.com/go-playground/validator/v10"
)

type TransactionService struct {
	Data productTransactions.ProductTransactionDataInterface
	validate *validator.Validate
}

func New(data productTransactions.ProductTransactionDataInterface) productTransactions.ProductTransactionServiceInterface {
	return &TransactionService{
		Data: data,
		validate: validator.New(),
	}
}

func (st *TransactionService) Create(transactionEntity productTransactions.ProductTransactionEntity) (productTransactions.ProductTransactionEntity, error) {
	st.validate = validator.New()
	errValidate := st.validate.StructExcept()
}