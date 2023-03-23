package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/users"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service users.UserServiceInterface
}

func New(s users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	users, err := h.Service.GetUser(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", UserEntityToUserResponse(users)))
}

func (h *UserHandler) Update(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	var formInput UserRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_, err := h.Service.Update(uint(userId), UserRequestToUserEntity(formInput))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	user, _ := h.Service.GetUser(uint(userId))

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", UserEntityToUserResponse(user)))
}

func (h *UserHandler) Delete(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	if err := h.Service.Delete(uint(userId)); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}

func (h *UserHandler) UpdateToSeller(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	var formInput UserRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	users, err := h.Service.UpdateToSeller(uint(userId), UserRequestToUserEntity(formInput))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", UserEntityToUserResponse(users)))
}

func (h *UserHandler) UpdateToProfile(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	f, err := c.FormFile("photo_profile")
	if err != nil {
		return err
	}

	newUrlProfile, err1 := h.Service.UpdateToProfile(uint(userId), f)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	data := map[string]any{
		"photo_profile": newUrlProfile,
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("success update profile", data))
}
