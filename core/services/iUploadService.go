package services

type IUploadService interface {
	UploadImage(image string, imageName string) error
}
