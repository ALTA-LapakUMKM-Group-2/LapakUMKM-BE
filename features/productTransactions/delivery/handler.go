package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/productTransactions"
	"lapakUmkm/utils/helpers"
	"net/http"
	"strconv"
	"strings"

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

func (hf *TransactionHandler) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transactionEntity, err := hf.service.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}
	transactioResponse := TransactionEntityToTransactionResponse(transactionEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("feedbacks detail", transactioResponse))
}

func (hf *TransactionHandler) CallBackMidtrans(c echo.Context) error {
	var form helpers.ResponseFromCallbackMidtrans

	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	idString := strings.Split(form.OrderId, "-")
	orderId, _ := strconv.Atoi(idString[1])

	err := hf.service.CallBackMidtrans(uint(orderId), form.TransactionStatus)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", ListTransactionToTransactionResponse(nil)))
}