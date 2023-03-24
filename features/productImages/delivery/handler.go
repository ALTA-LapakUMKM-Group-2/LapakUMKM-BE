package delivery

import (
	"lapakUmkm/features/productImages"
	"lapakUmkm/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductImagesHandler struct {
	Service productImages.ProductServiceInterface
}

func New(srv productImages.ProductServiceInterface) *ProductImagesHandler {
	return &ProductImagesHandler{
		Service: srv,
	}
}

func (h *ProductImagesHandler) Create(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	f, err := c.FormFile("photo_product")
	if err != nil {
		return err
	}

	data, err1 := h.Service.Create(uint(id), f)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("success update profile", data))
}
