package delivery

import (
	// "lapakUmkm/app/middlewares"
	"lapakUmkm/features/carts"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type BasketHandler struct {
	srv carts.CartService
}

func New(srv carts.CartService) *BasketHandler {
	return &BasketHandler{
		srv: srv,
	}
}

func (bah *BasketHandler) Add(c echo.Context) error {
	var formInput CartRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	// claim := middlewares.ClaimsToken(c)
	// formInput.UserID = uint(claim.Id)
	newBasket := carts.Core{}
	copier.Copy(&newBasket, &formInput)
	data, err := bah.srv.Add(newBasket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}
	res := AddResponse{}
	copier.Copy(&res, &data)
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("success add product to cart", res))
}
