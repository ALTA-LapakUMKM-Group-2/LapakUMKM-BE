package delivery

import "lapakUmkm/features/productImages"

type ProductImagesResponse struct {
	Image string `json:"image"`
}

func ProductImagesEntityToProductImagesResponse(i productImages.ProductImagesEntity) ProductImagesResponse {
	return ProductImagesResponse{
		Image: "https://storage.googleapis.com/images_lapak_umkm/product/" + i.Image,
	}
}
