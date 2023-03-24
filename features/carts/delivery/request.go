package delivery

type NewCartRequest struct {
	UserId     uint	`json:"user_id"`
	ProductId  uint `json:"product_id" form:"product_id"`
	ProductPcs int  `json:"product_pcs" form:"product_pcs"`
}