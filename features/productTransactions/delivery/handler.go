package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/productTransactions"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	service productTransactions.ProductTransactionServiceInterface
}

func New(srv productTransactions.ProductTransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{
		service: srv,
	}
}

func (ht *TransactionHandler) Create(c echo.Context) error {
	var formInput TransactionRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	userId := middlewares.ClaimsToken(c).Id
	user := TransactionRequestToTransactionEntity(&formInput)
	user.UserId = uint(userId)

	transaction, err := ht.service.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", TransactionEntityToTransactionResponse(transaction)))
}