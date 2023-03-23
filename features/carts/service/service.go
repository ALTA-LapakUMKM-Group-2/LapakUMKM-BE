package service

import (
	"lapakUmkm/features/carts"

	"github.com/go-playground/validator/v10"
)

type CartService struct {
	data carts.CartData
	vld  *validator.Validate
}

func New(data carts.CartData) carts.CartService {
	return &CartService{
		data: data,
		vld:  validator.New(),
	}
}

// Add implements carts.CartService
func (bas *CartService) Add(newCart carts.Core) (carts.Core, error) {
	// Check input validation
	errVld := bas.vld.Struct(newCart)
	if errVld != nil {
		return carts.Core{}, errVld
	}
	tmp, err := bas.data.Add(newCart)
	if err != nil {
		return carts.Core{}, err
	}
	return tmp, nil
}
