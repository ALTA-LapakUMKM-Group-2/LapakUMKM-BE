package service

import (
	"lapakUmkm/features/productTransactions"
	"lapakUmkm/utils/helpers"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type transactionService struct {
	Data     productTransactions.ProductTransactionDataInterface
	validate *validator.Validate
}

func New(data productTransactions.ProductTransactionDataInterface) productTransactions.ProductTransactionServiceInterface {
	return &transactionService{
		Data:     data,
		validate: validator.New(),
	}
}

func (st *transactionService) Create(transactionEntity productTransactions.ProductTransactionEntity) (productTransactions.ProductTransactionEntity, error) {
	st.validate = validator.New()
	errValidate := st.validate.StructExcept(transactionEntity, "User", "Product")
	if errValidate != nil {
		return productTransactions.ProductTransactionEntity{}, errValidate
	}
	transactionId, err := st.Data.Store(transactionEntity)
	if err != nil {
		return productTransactions.ProductTransactionEntity{}, err
	}

	totalProduct := transactionEntity.TotalProduct
	totalPayment := transactionEntity.TotalPayment

	//call midtrans
	postData := map[string]any{
		"order_id":  "lapakumkm-" + strconv.Itoa(int(transactionId)),
		"nominal":   totalPayment,
		"firstname": "LapakUMKM",
		"lastname":  "Product",
		"email":     "email" + strconv.Itoa(int(transactionId)) + "@gmail.com",
		"phone":     "000",
	}

	paymentLink, err1 := helpers.PostMidtrans(postData)
	if err1 != nil {
		return productTransactions.ProductTransactionEntity{}, err
	} else {
		//midtrans
		update := productTransactions.ProductTransactionEntity{
			TotalProduct:  totalProduct,
			TotalPayment:  totalPayment,
			PaymentStatus: "pending",
			PaymentLink:   paymentLink,
		}
		//if ok
		st.Data.Edit(update, transactionId)
	}

	return st.Data.SelectById(transactionId)
}

func (st *transactionService) MyTransactionHistory(myId, userId uint) ([]productTransactions.ProductTransactionEntity, error) {
	return st.Data.SelectAll(userId)
}

func (st *transactionService) GetById(id uint) (productTransactions.ProductTransactionEntity, error) {
	return st.Data.SelectById(id)
}

func (st *transactionService) CallBackMidtrans(id uint, status string) error {
	transactions := productTransactions.ProductTransactionEntity{
		PaymentStatus: status,
	}
	err := st.Data.Edit(transactions, id)
	if err != nil {
		return err
	}
	return nil
}
