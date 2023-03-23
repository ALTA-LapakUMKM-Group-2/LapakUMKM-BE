package delivery

import (
	"lapakUmkm/features/discussions"
	// "lapakUmkm/app/middlewares"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DiscussionHandler struct {
	Service discussions.DiscussionServiceInterface
}

func New(srv discussions.DiscussionServiceInterface) *DiscussionHandler{
	return &DiscussionHandler{
		Service: srv,
	}
}

func (hd *DiscussionHandler) Create(c echo.Context) error{
	var formInput DiscussionRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	// claim := middlewares.ClaimsToken(c)
	// user_id := claim.Id
	// formInput.UserId = uint(user_id)

	feedback, err := hd.Service.Create(DiscussionRequestToDiscussionEntity(&formInput))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", DiscussionEntityToDiscussionResponse(feedback)))
}