package productsimages

import (
	"time"
)

type ProductImages struct {
	Id        uint
	ProductId uint
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
