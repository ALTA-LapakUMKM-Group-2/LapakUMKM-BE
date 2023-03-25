package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/discussions"
	"lapakUmkm/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DiscussionHandler struct {
	Service discussions.DiscussionServiceInterface
}

func New(srv discussions.DiscussionServiceInterface) *DiscussionHandler {
	return &DiscussionHandler{
		Service: srv,
	}
}

func (hd *DiscussionHandler) Create(c echo.Context) error {
	var formInput DiscussionRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	userId := middlewares.ClaimsToken(c).Id
	user := DiscussionRequestToDiscussionEntity(&formInput)
	user.UserId = uint(userId)

	discussion, err := hd.Service.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", DiscussionEntityToDiscussionResponse(discussion)))
}

func (hd *DiscussionHandler) Update(c echo.Context) error {
	var formInput DiscussionPutRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	userId := middlewares.ClaimsToken(c).Id

	discussion, err := hd.Service.Update(DiscussionPutRequestToDiscussionEntity(&formInput), uint(id), uint(userId))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", DiscussionEntityToDiscussionResponse(discussion)))
}