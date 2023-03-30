package delivery

import (
	"lapakUmkm/features/productTransactionDetails"
	"lapakUmkm/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductTransactionDetailHandler struct {
	Service productTransactionDetails.ProductTransactionDetailServiceInterface
}

func New(srv productTransactionDetails.ProductTransactionDetailServiceInterface) *ProductTransactionDetailHandler {
	return &ProductTransactionDetailHandler{
		Service: srv,
	}
}

func (h *ProductTransactionDetailHandler) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	productDetail, err := h.Service.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", EntityToResponse(productDetail)))
}

func (h *ProductTransactionDetailHandler) GetByTransaksiId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	productDetail, _ := h.Service.GetByTransaksiId(uint(id))
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", ListEntityToResponse(productDetail)))
}

func (h *ProductTransactionDetailHandler) Create(c echo.Context) error {
	var formInput ProductTransactionDetailsRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	productDetails, err := h.Service.Create(RequestToEntity(&formInput))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", EntityToResponse(productDetails)))
}
