package categories

import (
	"time"
)

type CategoryEntity struct {
	Id        uint
	Category  string `validate:"required,min=3,max=20"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryServiceInterface interface {
	GetAll() ([]CategoryEntity, error)
	GetById(id uint) (CategoryEntity, error)
	Create(categoryEntity CategoryEntity) (CategoryEntity, error)
	Update(categoryEntity CategoryEntity, id uint) (CategoryEntity, error)
	Delete(id uint) error
}

type CategoryDataInterface interface {
	SelectAll() ([]CategoryEntity, error)
	SelectById(id uint) (CategoryEntity, error)
	Store(categoryEntity CategoryEntity) (uint, error)
	Edit(categoryEntity CategoryEntity, id uint) error
	Destroy(id uint) error
}
