package data

import (
	"lapakUmkm/features/products"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) products.ProductDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll(productFilter products.ProductFilter) ([]products.ProductEntity, error) {
	var products []Product

	query := q.db.Preload("User").Preload("Category")
	if productFilter.PriceMin != 0 {
		query.Where("products.price >= ?", productFilter.PriceMin)
	}

	if productFilter.PriceMax != 0 {
		query.Where("products.price <= ?", productFilter.PriceMax)
	}

	if productFilter.CategoryId != 0 {
		query.Where("products.category_id = ?", productFilter.CategoryId)
	}

	if err := query.Find(&products); err.Error != nil {
		return nil, err.Error
	}
	return ListProductToProductEntity(products), nil
}

func (q *query) SelectById(id uint) (products.ProductEntity, error) {
	var product Product
	if err := q.db.Preload("User").Preload("Category").First(&product, id); err.Error != nil {
		return products.ProductEntity{}, err.Error
	}
	return ProductToProductEntity(product), nil
}

func (q *query) Store(productEntity products.ProductEntity) (uint, error) {
	product := ProductEntityToProduct(productEntity)
	if err := q.db.Create(&product); err.Error != nil {
		return 0, err.Error
	}
	return product.ID, nil
}

func (q *query) Edit(productEntity products.ProductEntity, id uint) error {
	product := ProductEntityToProduct(productEntity)
	if err := q.db.Where("id", id).Updates(&product); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var product Product
	if err := q.db.Delete(&product, id); err.Error != nil {
		return err.Error
	}
	return nil
}
