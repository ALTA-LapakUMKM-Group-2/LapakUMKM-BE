package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/productTransactions"
	"lapakUmkm/utils/helpers"
	"net/http"
	"strconv"

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

func (ht *TransactionHandler) MyTransactionHistory(c echo.Context) error {
	myId := middlewares.ClaimsToken(c).Id
	userId, _ := strconv.Atoi(c.Param("id"))
	feedbackEntity, err := ht.service.MyTransactionHistory(uint(userId), uint(myId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listFeedbackResponse := ListTransactionToTransactionResponse(feedbackEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("all your feedbacks", listFeedbackResponse))
}
