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
	ProductPcs   int
	ProductPrice int64
	ProductImage string
	LapakName    string
	LapakAddress string
}

func CoreToCart(data carts.Core) Cart {
	return Cart{
		Model:        gorm.Model{ID: data.Id},
		UserId:       data.UserId,
		ProductId:    data.ProductId,
		ProductName:  data.ProductName,
		ProductPcs:   data.ProductPcs,
		ProductPrice: data.ProductPrice,
		ProductImage: data.ProductImage,
		LapakName:    data.LapakName,
		LapakAddress: data.LapakAddress,
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
		ProductImage: data.ProductImage,
		LapakName:    data.LapakName,
		LapakAddress: data.LapakAddress,
	}
}
