package delivery

type ProductImagesResponse struct {
	Image string `json:"image"`
}

func ProductImagesEntityToProductImagesResponse() ProductImagesResponse {
	return ProductImagesResponse{}
}
