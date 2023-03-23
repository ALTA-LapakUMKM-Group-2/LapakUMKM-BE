package delivery

type AddResponse struct {
	Id           uint   `json:"id"`
	ProductId    uint   `json:"product_id"`
	ProductPcs   int    `json:"product_pcs"`
}

type GetResponse struct {
	Id           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	ProductId    uint   `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPcs   int    `json:"product_pcs"`
	ProductPrice int64  `json:"product_price"`
	ProductImage string `json:"product_image"`
	LapakName    string `json:"lapak_name"`
	LapakAddress string `json:"lapak_address"`
}
