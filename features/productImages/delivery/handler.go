package delivery

import (
	"lapakUmkm/features/productImages"
	"lapakUmkm/features/products"
	"lapakUmkm/utils/helpers"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductImagesHandler struct {
	Service        productImages.ProductServiceInterface
	ServiceProduct products.ProductServiceInterface
}

func New(srv productImages.ProductServiceInterface, srv2 products.ProductServiceInterface) *ProductImagesHandler {
	return &ProductImagesHandler{
		Service:        srv,
		ServiceProduct: srv2,
	}
}

func (h *ProductImagesHandler) Create(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	_, err1 := h.ServiceProduct.GetById(uint(productId))
	if err1 != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err1.Error()))
	}

	const maxFileSize = 1024 * 1024

	file, err := c.FormFile("photo_product")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	size := file.Size
	if size > maxFileSize {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("file too large (max 1 MB)"))
	}

	fileExtension := filepath.Ext(file.Filename)
	fileExtension = strings.ToLower(fileExtension)

	if fileExtension != ".jpg" && fileExtension != ".png" && fileExtension != ".jpeg" {
		return c.JSON(http.StatusUnsupportedMediaType, helpers.ResponseFail("only image extention (png,jpg, or jpeg)"))
	}

	data2, err1 := h.Service.GetByProductId(uint(productId))
	if err1 != nil {
		if len(data2) >= 5 {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail("max 5 images for 1 product"))
		}

		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	data, err1 := h.Service.Create(uint(productId), file)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("success add image", ProductImagesEntityToProductImagesResponse(data)))
}

func (h *ProductImagesHandler) Delete(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	_, err1 := h.ServiceProduct.GetById(uint(productId))
	if err1 != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err1.Error()))
	}

	id, _ := strconv.Atoi(c.Param("photo_id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("success delete image", nil))
}

func (h *ProductImagesHandler) GetByProductId(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	_, err1 := h.ServiceProduct.GetById(uint(productId))
	if err1 != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err1.Error()))
	}

	data, err := h.Service.GetByProductId(uint(productId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", ListResponse(data)))
}
