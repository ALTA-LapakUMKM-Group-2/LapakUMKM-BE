package helpers

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadPhotoProfile(file multipart.File, fileNameInGCS string) error {
	bucketName := "images_lapak_umkm"
	pathFile := "profile/" + fileNameInGCS

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("filecred.json"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(pathFile)
	writer := obj.NewWriter(ctx)

	if _, err := io.Copy(writer, file); err != nil {
		fmt.Println(err)
		return err
	}

	if err := writer.Close(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeletePhotoProfile(fileName string) error {
	bucketName := "images_lapak_umkm"
	pathFile := "profile/" + fileName

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("filecred.json"))
	if err != nil {
		return err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	image := bucket.Object(pathFile)

	if err := image.Delete(ctx); err != nil {
		return err
	}

	return nil
}

func UploadPhotoProduct(file multipart.File, fileNameInGCS string) error {
	bucketName := "images_lapak_umkm"
	pathFile := "products/" + fileNameInGCS

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("filecred.json"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(pathFile)
	writer := obj.NewWriter(ctx)

	if _, err := io.Copy(writer, file); err != nil {
		fmt.Println(err)
		return err
	}

	if err := writer.Close(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeletePhotoProduct(fileName string) error {
	bucketName := "images_lapak_umkm"
	pathFile := "profile/" + fileName

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("filecred.json"))
	if err != nil {
		return err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	image := bucket.Object(pathFile)

	if err := image.Delete(ctx); err != nil {
		return err
	}

	return nil
}
