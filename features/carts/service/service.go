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

// Update implements carts.CartService
func (cs *CartService) Update(updateCart carts.Core) (carts.Core, error) {
	tmp, err := cs.data.Update(updateCart)
	if err != nil {
		return carts.Core{}, err
	}
	return tmp, nil
}

// Delete implements carts.CartService
func (cs *CartService) Delete(userID, cartID uint) error {
	err := cs.data.Delete(userID, cartID)
	if err != nil {
		return err
	}
	return nil
}

// CartByID implements carts.CartService
func (cs *CartService) CartByID(userID uint, cart []uint) ([]carts.Core, error) {
	tmp, err := cs.data.CartByID(userID, cart)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

// BuyNow implements carts.CartService
func (cs *CartService) BuyNow(buyNow carts.Core) (carts.Core, error) {
	errVld := cs.vld.Struct(buyNow)
	if errVld != nil {
		return carts.Core{}, errVld
	}
	tmp, err := cs.data.BuyNow(buyNow)
	if err != nil {
		return carts.Core{}, err
	}
	return tmp, nil
}