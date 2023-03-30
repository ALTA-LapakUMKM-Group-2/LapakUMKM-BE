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

func (hd *DiscussionHandler) Delete(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	id, _ := strconv.Atoi(c.Param("id"))

	if err := hd.Service.Delete(uint(id), uint(userId)); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}

func (hd *DiscussionHandler) GetDiscussionByProductId(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	discussions, err := hd.Service.GetDiscussionByProductId(uint(productId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listDiscussionsResponse := ListDiscussionEntityToDiscussionResponse(discussions)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("discussion by product id", listDiscussionsResponse))
}

func (hd *DiscussionHandler) GetAll(c echo.Context) error {
	myId := middlewares.ClaimsToken(c).Id
	userId, _ := strconv.Atoi(c.Param("id"))
	feedbackEntity, err := hd.Service.GetAll(uint(userId), uint(myId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listDiscussionsResponse := ListDiscussionEntityToDiscussionResponse(feedbackEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("all your discussions", listDiscussionsResponse))
}

func (hd *DiscussionHandler) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	discussionEntity, err := hd.Service.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}
	discussionResponse := DiscussionEntityToDiscussionResponse(discussionEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("discussion detail", discussionResponse))
}
