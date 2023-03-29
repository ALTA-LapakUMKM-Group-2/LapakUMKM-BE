package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/productImages"
	"lapakUmkm/features/products"
	"lapakUmkm/utils/helpers"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service       products.ProductServiceInterface
	ServiceImages productImages.ProductServiceInterface
}

func New(srv products.ProductServiceInterface, srv2 productImages.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		Service:       srv,
		ServiceImages: srv2,
	}
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	priceMin, _ := strconv.Atoi(c.QueryParam("price_min"))
	priceMax, _ := strconv.Atoi(c.QueryParam("price_max"))
	rating, _ := strconv.Atoi(c.QueryParam("rating"))
	categoryId, _ := strconv.Atoi(c.QueryParam("category_id"))
	userId, _ := strconv.Atoi(c.QueryParam("user_id"))

	productFilter := products.ProductFilter{
		PriceMin:   priceMin,
		PriceMax:   priceMax,
		Rating:     float64(rating),
		CategoryId: uint(categoryId),
		UserId:     uint(userId),
	}

	products, err := h.Service.GetAll(productFilter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data null"))
	}

	listProductsResponse := ListProductEntityToProductResponse(products)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", listProductsResponse))
}

func (h *ProductHandler) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	productEntity, err := h.Service.GetById(uint(id))
	if err != nil {
		str := err.Error()
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(str))
	}
	productResponse := ProductEntityToProductResponse(productEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", productResponse))
}

func (h *ProductHandler) Create(c echo.Context) error {
	var formInput ProductRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	userId := middlewares.ClaimsToken(c).Id
	productEntity := ProductRequestToProductEntity(&formInput)
	productEntity.UserId = uint(userId)

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

	product, err := h.Service.Create(productEntity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	_, err1 := h.ServiceImages.Create(product.Id, file)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	products, _ := h.Service.GetById(product.Id)

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", ProductEntityToProductResponse(products)))
}

func (h *ProductHandler) Update(c echo.Context) error {
	var formInput ProductRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	userId := middlewares.ClaimsToken(c).Id

	//check data exist
	checkDataExist, errExist := h.Service.GetById(uint(id))
	if errExist != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(errExist.Error()))
	}

	if checkDataExist.UserId != uint(userId) {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("can't update this product id"))
	}

	product, err := h.Service.Update(ProductRequestToProductEntity(&formInput), uint(id), uint(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", ProductEntityToProductResponse(product)))
}

func (t *ProductHandler) Delete(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	id, _ := strconv.Atoi(c.Param("id"))

	if err := t.Service.Delete(uint(id), uint(userId)); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}
