package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/products"
	"lapakUmkm/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service products.ProductServiceInterface
}

func New(srv products.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		Service: srv,
	}
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	priceMin, _ := strconv.Atoi(c.QueryParam("price_min"))
	priceMax, _ := strconv.Atoi(c.QueryParam("price_max"))
	rating, _ := strconv.Atoi(c.QueryParam("rating"))
	categoryId, _ := strconv.Atoi(c.QueryParam("category_id"))
	userId, _ := strconv.Atoi(c.QueryParam("userId"))

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
	userEntity := ProductRequestToProductEntity(&formInput)
	userEntity.UserId = uint(userId)

	team, err := h.Service.Create(userEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", ProductEntityToProductResponse(team)))
}

func (h *ProductHandler) Update(c echo.Context) error {
	var formInput ProductRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	userId := middlewares.ClaimsToken(c).Id

	product, err := h.Service.Update(ProductRequestToProductEntity(&formInput), uint(id), uint(userId))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
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
