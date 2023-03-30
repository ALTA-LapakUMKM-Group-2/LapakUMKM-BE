package data

import (
	"lapakUmkm/features/productTransactionDetails"
	"lapakUmkm/features/products"
	product "lapakUmkm/features/products/data"
	"reflect"
)

type ProductTransactionDetail struct {
	Id                   uint
	ProductTransactionID uint
	ProductId            uint
	Product              product.Product
	TotalProduct         int
	Rating               float64
}

func EntityToModel(p productTransactionDetails.ProductTransactionDetailEntity) ProductTransactionDetail {
	return ProductTransactionDetail{
		ProductTransactionID: p.ProductTransactionID,
		ProductId:            p.ProductId,
		TotalProduct:         p.TotalProduct,
	}
}

func ModelToEntity(p ProductTransactionDetail) productTransactionDetails.ProductTransactionDetailEntity {
	result := productTransactionDetails.ProductTransactionDetailEntity{
		Id:                   p.Id,
		ProductTransactionID: p.ProductTransactionID,
		ProductId:            p.ProductId,
		TotalProduct:         p.TotalProduct,
		Rating:               p.Rating,
	}

	if !reflect.ValueOf(p.Product).IsZero() {
		result.Product = products.ProductEntity{
			ProductName: p.Product.ProductName,
			Price:       p.Product.Price,
		}
	}

	return result
}

func ListModelToEntity(p []ProductTransactionDetail) []productTransactionDetails.ProductTransactionDetailEntity {
	var teamEntity []productTransactionDetails.ProductTransactionDetailEntity
	for _, v := range p {
		teamEntity = append(teamEntity, ModelToEntity(v))
	}
	return teamEntity
}
