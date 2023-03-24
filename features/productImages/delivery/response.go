package delivery

import "lapakUmkm/features/productImages"

type ProductImagesResponse struct {
	Id    uint   `json:"id"`
	Image string `json:"image"`
}

func ProductImagesEntityToProductImagesResponse(i productImages.ProductImagesEntity) ProductImagesResponse {
	return ProductImagesResponse{
		Id:    i.Id,
		Image: "https://storage.googleapis.com/images_lapak_umkm/product/" + i.Image,
	}
}

func ListResponse(data []productImages.ProductImagesEntity) []ProductImagesResponse {
	var response []ProductImagesResponse
	for _, v := range data {
		response = append(response, ProductImagesEntityToProductImagesResponse(v))
	}

	return response
}
