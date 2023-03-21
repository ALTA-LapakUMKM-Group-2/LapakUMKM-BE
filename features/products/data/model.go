package data

import (
	dataSubCategories "lapakUmkm/features/subcategories/data"
	dataUser "lapakUmkm/features/users/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserId        uint
	User          *dataUser.User `gorm:"foreignKey:UserId"`
	SubCategoryId uint
	SubCategory   *dataSubCategories.SubCategory `gorm:"foreignKey:SubCategoryId"`
	ProductName   string
	Description   string
	Price         int
}
