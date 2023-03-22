package delivery

import (
	"lapakUmkm/features/categories"
	"lapakUmkm/utils/helpers"
	"net/http"

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
	listTeamResponse := ListCategoryEntityToCategoryResponse(categoryEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", listTeamResponse))
}
