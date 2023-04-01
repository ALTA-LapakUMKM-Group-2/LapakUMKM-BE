package carts

type Core struct {
	Id           uint
	UserId       uint
	ProductId    uint `validate:"required"`
	ProductName  string
	ProductPcs   int64 `validate:"required"`
	ProductPrice int64
	SubTotal     int64
	ProductImage string
	LapakName    string
	LapakAddress string
	PhotoProfile string
}

type CartService interface {
	Add(cart Core) (Core, error)
	MyCart(userID uint) ([]Core, error)
	Update(updateCart Core) (Core, error)
	Delete(userID, cartID uint) error
}

type CartData interface {
	Add(newCart Core) (Core, error)
	MyCart(userID uint) ([]Core, error)
	Update(updateCart Core) (Core, error)
	Delete(userID, cartID uint) error
}