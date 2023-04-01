package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/chats"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	service chats.ServiceInterface
}

func New(srv chats.ServiceInterface) *ChatHandler {
	return &ChatHandler{
		service: srv,
	}
}

func (h *ChatHandler) Create(c echo.Context) error {
	var formInput ChatRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	senderId := middlewares.ClaimsToken(c).Id
	user := ReqToEntity(&formInput)
	user.SenderId = uint(senderId)

	chat, err := h.service.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", EntityToResponse(chat)))
}

func (h *ChatHandler) GetByRoomId(c echo.Context) error {
	roomId := (c.Param("id"))
	chatss, err := h.service.GetByRoomId(roomId)

	if len(chatss) == 0 {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("messages not found"))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listChatsResponse := ListEntityToResponse(chatss)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("chats by room id", listChatsResponse))
}
