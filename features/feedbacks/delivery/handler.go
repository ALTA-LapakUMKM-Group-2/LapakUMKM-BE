package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/feedbacks"
	"lapakUmkm/utils/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FeedbackHandler struct {
	service feedbacks.FeedbackServiceInterface
}

func New(srv feedbacks.FeedbackServiceInterface) *FeedbackHandler {
	return &FeedbackHandler{
		service: srv,
	}
}

func (hf *FeedbackHandler) Create(c echo.Context) error {
	var formInput FeedbackRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	userId := middlewares.ClaimsToken(c).Id
	user := FeedbackRequestToFeedbackEntity(&formInput)
	user.UserId = uint(userId)

	feedback, err := hf.service.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", FeedbackEntityToFeedbackPostResponse(feedback)))
}

func (hf *FeedbackHandler) Update(c echo.Context) error {
	var formInput FeedbackPutRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	id, _ := strconv.Atoi(c.Param("id"))
	userId := middlewares.ClaimsToken(c).Id

	feedback, err := hf.service.Update(FeedbackPutRequestToFeedbackEntity(&formInput), uint(id), uint(userId))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", FeedbackEntityToFeedbackPutResponse(feedback)))
}

func (hf *FeedbackHandler) Delete(c echo.Context) error {
	userId := middlewares.ClaimsToken(c).Id
	id, _ := strconv.Atoi(c.Param("id"))

	if err := hf.service.Delete(uint(id), uint(userId)); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}

func (hf *FeedbackHandler) GetFeedbackByProductId(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	feedbacks, err := hf.service.GetFeedbackByProductId(uint(productId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listFeedbacksResponse := ListFeedbackToFeedbackGetResponse(feedbacks)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("feedback by product id", listFeedbacksResponse))
}

func (hf *FeedbackHandler) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	feedbackEntity, err := hf.service.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}
	feedbackResponse := FeedbackEntityToFeedbackGetResponse(feedbackEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("feedbacks detail", feedbackResponse))
}
