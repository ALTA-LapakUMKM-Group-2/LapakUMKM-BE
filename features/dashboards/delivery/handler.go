package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/dashboards"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DiscussionHandler struct {
	Service dashboards.DashboardServiceInterface
}

func New(srv dashboards.DashboardServiceInterface) *DiscussionHandler {
	return &DiscussionHandler{
		Service: srv,
	}
}

func (h *DiscussionHandler) Get(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	data, _ := h.Service.GetByUserId(uint(userId))
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Dashboard Success", EntityToResponse(data)))
}

func (h *DiscussionHandler) Create(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	h.Service.Create(uint(userId))
	h.Service.UpdateData(uint(userId))
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Dashboard Success", nil))
}
