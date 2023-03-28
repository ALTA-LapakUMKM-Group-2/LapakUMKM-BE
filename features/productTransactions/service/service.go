package service

import (
	"fmt"
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
	fmt.Println("1")
	st.validate = validator.New()
	errValidate := st.validate.StructExcept(transactionEntity, "User")
	if errValidate != nil {
		return productTransactions.ProductTransactionEntity{}, errValidate
	}

	fmt.Println("2")
	transactionId, err := st.Data.Store(transactionEntity)
	if err != nil {
		return productTransactions.ProductTransactionEntity{}, err
	}

	totalProduct := transactionEntity.TotalProduct
	totalPayment := transactionEntity.TotalPayment

	//call midtrans
	orderId := "lapakumkm-" + strconv.Itoa(int(transactionId))
	postData := map[string]any{
		"order_id":  orderId,
		"nominal":   totalPayment,
		"firstname": "LapakUMKM",
		"lastname":  "Product",
		"email":     "email" + strconv.Itoa(int(transactionId)) + "@gmail.com",
		"phone":     "000",
	}

	paymentLink, err1 := helpers.PostMidtrans(postData)
	if err1 != nil {
		fmt.Println("3")
		return productTransactions.ProductTransactionEntity{}, err
	} else {
		fmt.Println("4")
		update := productTransactions.ProductTransactionEntity{
			TotalProduct:  totalProduct,
			TotalPayment:  totalPayment,
			PaymentStatus: "pending",
			PaymentLink:   paymentLink,
			OrderId:       orderId,
		}
		//if ok
		st.Data.Edit(update, transactionId)
	}

	fmt.Println("5")

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
