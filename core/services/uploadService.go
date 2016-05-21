package services

import (
	"encoding/base64"
	"os"
	"strings"
	"io"
)

type IUploadService interface{
	 UploadImage(image string, imageName string) error 
}

type uploadService struct{
	StoreLocation string
}

func UploadService(location string) uploadService{
	service := uploadService{}
	service.StoreLocation = location
	
	return service
}

func (service uploadService) UploadImage(image string, imageName string) error {
	plainBase64 := strings.Replace(image, "data:image/png;base64,", "", 1)
	
	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(plainBase64))
	
	img, err := os.Create(service.StoreLocation + imageName)
	
	if(err == nil){
		defer img.Close()
		_, err = io.Copy(img, imageReader)
	}
	
	return err
}
