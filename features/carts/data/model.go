package data

import (
	"lapakUmkm/features/carts"
	dataProduct "lapakUmkm/features/products/data"
	dataUser "lapakUmkm/features/users/data"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId       uint
	User         *dataUser.User `gorm:"foreignKey:UserId"`
	ProductId    uint
	Product      *dataProduct.Product `gorm:"foreignKey:ProductId"`
	ProductName  string
	ProductPcs   int64
	ProductPrice int64
	SubTotal     int64
	ProductImage string
	LapakName    string
	LapakAddress string
	PhotoProfile string
}

func CoreToCart(data carts.Core) Cart {
	return Cart{
		Model:        gorm.Model{ID: data.Id},
		UserId:       data.UserId,
		ProductId:    data.ProductId,
		ProductName:  data.ProductName,
		ProductPcs:   data.ProductPcs,
		ProductPrice: data.ProductPrice,
		SubTotal:     data.SubTotal,
		ProductImage: data.ProductImage,
		LapakName:    data.LapakName,
		LapakAddress: data.LapakAddress,
		PhotoProfile: data.PhotoProfile,
	}
}

func CartToCore(data Cart) carts.Core {
	return carts.Core{
		Id:           data.ID,
		UserId:       data.UserId,
		ProductId:    data.ProductId,
		ProductName:  data.ProductName,
		ProductPcs:   data.ProductPcs,
		ProductPrice: data.ProductPrice,
		SubTotal:     data.SubTotal,
		ProductImage: data.ProductImage,
		LapakName:    data.LapakName,
		LapakAddress: data.LapakAddress,
		PhotoProfile: data.PhotoProfile,
	}
}

func ListCartToCore(data []Cart) []carts.Core {
	var dataCore []carts.Core
	for _, v := range data {
		dataCore = append(dataCore, CartToCore(v))
	}
	return dataCore
}
