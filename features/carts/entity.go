package carts

type Core struct {
	Id           uint
	UserId       uint
	ProductId    uint `validate:"required"`
	ProductName  string
	ProductPcs   int `validate:"required"`
	ProductPrice int64
	ProductImage string
	LapakName    string
	LapakAddress string
}

type CartService interface {
	Add(newCart Core) (Core, error)
	MyCart(userID uint) ([]Core, error)
}

type CartData interface {
	Add(newCart Core) (Core, error)
	MyCart(userID uint) ([]Core, error)
}