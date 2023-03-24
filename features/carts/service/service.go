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
func (cs *CartService) Add(newCart carts.Core) (carts.Core, error) {
	errVld := cs.vld.Struct(newCart)
	if errVld != nil {
		return carts.Core{}, errVld
	}
	tmp, err := cs.data.Add(newCart)
	if err != nil {
		return carts.Core{}, err
	}
	return tmp, nil
}

// MyCart implements carts.CartService
func (cs *CartService) MyCart(userID uint) ([]carts.Core, error) {
	tmp, err := cs.data.MyCart(userID)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}