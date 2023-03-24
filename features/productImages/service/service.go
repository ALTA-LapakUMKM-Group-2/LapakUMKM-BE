package service

import (
	"lapakUmkm/features/productImages"
	"lapakUmkm/utils/helpers"
	"mime/multipart"
	"strconv"
	"time"
)

type ProductImagesService struct {
	Data productImages.ProductDataInterface
}

func New(data productImages.ProductDataInterface) productImages.ProductServiceInterface {
	return &ProductImagesService{
		Data: data,
	}
}

func (s *ProductImagesService) Create(productId uint, file *multipart.FileHeader) (productImages.ProductImagesEntity, error) {
	var empty productImages.ProductImagesEntity
	blobFile, err := file.Open()
	if err != nil {
		return empty, err
	}

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	newFileName := timestamp + "_" + strconv.Itoa(int(productId)) + ".png"
	helpers.UploadPhotoProduct(blobFile, newFileName)

	var request = productImages.ProductImagesEntity{
		ProductId: productId,
		Image:     newFileName,
	}

	id, err := s.Data.Store(request)
	if err != nil {
		return empty, err
	}

	return s.Data.SelectById(id)
}

func (s *ProductImagesService) Delete(id uint) error {
	dataImage, err := s.Data.SelectById(id)
	if err != nil {
		return err
	}

	if err := helpers.DeletePhotoProduct(dataImage.Image); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}

// GetByProductId implements productImages.ProductServiceInterface
func (s *ProductImagesService) GetByProductId(productId uint) ([]productImages.ProductImagesEntity, error) {
	return s.Data.SelectByProductId(productId)
}
