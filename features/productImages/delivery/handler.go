package delivery

import (
	"errors"
	"lapakUmkm/features/productImages"
	"lapakUmkm/utils/helpers"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

	const maxFileSize = 1024 * 1024
	err1 := c.Request().ParseMultipartForm(maxFileSize)
	if err1 != nil {
		return errors.New("file to large, max 1 MB")
	}

	file, err := c.FormFile("photo_product")
	if err != nil {
		return err
	}

	fileExtension := filepath.Ext(file.Filename)
	fileExtension = strings.ToLower(fileExtension)

	if fileExtension == ".jpg" || fileExtension == ".png" || fileExtension == ".jpeg" {
		return nil
	}

	data, err1 := h.Service.Create(uint(id), file)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("success add image", ProductImagesEntityToProductImagesResponse(data)))
}

func (h *ProductImagesHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("photo_id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("success delete image", nil))
}

func (h *ProductImagesHandler) GetByProductId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.Service.GetByProductId(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", ListResponse(data)))
}
