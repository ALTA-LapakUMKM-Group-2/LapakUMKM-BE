package delivery

import (
	"lapakUmkm/features/categories"
	"lapakUmkm/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Service categories.CategoryServiceInterface
}

func New(srv categories.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{
		Service: srv,
	}
}

func (h *CategoryHandler) GetAll(c echo.Context) error {
	categoryEntity, err := h.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	dataCategory := ListCategoryEntityToCategoryResponse(categoryEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", dataCategory))
}

func (t *CategoryHandler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	teamEntity, err := t.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}

	dataCategory := CategoryEntityToCategoryResponse(teamEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", dataCategory))
}

func (t *CategoryHandler) Create(c echo.Context) error {
	var formInput CategoryRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	team, err := t.Service.Create(CategoryRequestToCategoryEntity(&formInput))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", CategoryEntityToCategoryResponse(team)))
}

func (t *CategoryHandler) Update(c echo.Context) error {
	var formInput CategoryRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	dataCategory, err := t.Service.Update(CategoryRequestToCategoryEntity(&formInput), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", CategoryEntityToCategoryResponse(dataCategory)))
}

func (t *CategoryHandler) Delete(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	if err := t.Service.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}
