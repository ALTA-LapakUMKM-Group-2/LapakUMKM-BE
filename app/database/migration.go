package database

import (
	"fmt"
	baskets "lapakUmkm/features/baskets/data"
	categories "lapakUmkm/features/categories/data"
	discussions "lapakUmkm/features/discussions/data"
	feedbacks "lapakUmkm/features/feedbacks/data"
	productTransactions "lapakUmkm/features/productTransactions/data"
	productTypes "lapakUmkm/features/productTypes/data"
	products "lapakUmkm/features/products/data"
	subcategories "lapakUmkm/features/subcategories/data"
	user "lapakUmkm/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&categories.Category{},
		&subcategories.SubCategory{},
		&products.Product{},
		&productTypes.ProductType{},
		&productTransactions.ProductTransaction{},
		&baskets.Basket{},
		&feedbacks.Feedback{},
		&discussions.Discussion{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
