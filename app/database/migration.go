package database

import (
	"fmt"
	carts "lapakUmkm/features/carts/data"
	categories "lapakUmkm/features/categories/data"
	discussions "lapakUmkm/features/discussions/data"
	feedbacks "lapakUmkm/features/feedbacks/data"
	productImages "lapakUmkm/features/productImages/data"
	productTransactionDetails "lapakUmkm/features/productTransactionDetails/data"
	productTransactions "lapakUmkm/features/productTransactions/data"
	products "lapakUmkm/features/products/data"
	user "lapakUmkm/features/users/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&categories.Category{},
		&products.Product{},
		&productTransactions.ProductTransaction{},
		&carts.Cart{},
		&feedbacks.Feedback{},
		&discussions.Discussion{},
		&productImages.ProductImages{},
		&productTransactionDetails.ProductTransactionDetail{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
