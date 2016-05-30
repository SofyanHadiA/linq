package products

import (
	"errors"

	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/core/services"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type ProductService struct {
	repo          core.IRepository
	uploadService services.IUploadService
}

func NewProductService(repo core.IRepository, uploadService services.IUploadService) ProductService {
	return ProductService{
		repo:          repo,
		uploadService: uploadService,
	}
}

func (service ProductService) CountAll() (int, error) {
	return service.repo.CountAll()
}

func (service ProductService) IsExist(id uuid.UUID) (bool, error) {
	return service.repo.IsExist(id)
}

func (service ProductService) GetAll(paging utils.Paging) (core.IModels, error) {
	return service.repo.GetAll(paging)
}

func (service ProductService) Get(id uuid.UUID) (core.IModel, error) {
	return service.repo.Get(id)
}

func (service ProductService) Create(model core.IModel) error {
	return service.repo.Insert(model)
}

func (service ProductService) Modify(model core.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		return service.repo.Update(model)
	} else {
		return productNotFoundError(model.GetId())
	}
}

func (service ProductService) UpdateProductPhoto(model core.IModel, imageString string) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		fileName := model.GetId().String() + ".png"
		err := service.uploadService.UploadImage(imageString, fileName)

		if err == nil {
			product, _ := model.(*Product)

			product.Image.String = fileName
			product.Image.Valid = true

			productRepo := service.repo.(productRepository)
			err = productRepo.UpdateProductPhoto(product)
		}

		return err
	} else {
		return productNotFoundError(model.GetId())
	}
}

func (service ProductService) Remove(model core.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		err := service.repo.Delete(model)
		return err
	} else {
		return productNotFoundError(model.GetId())
	}
}

func (service ProductService) RemoveBulk(productIds []uuid.UUID) error {
	for _, id := range productIds {
		if exist, _ := service.repo.IsExist(id); !exist {
			return productNotFoundError(id)
		}
	}

	err := service.repo.DeleteBulk(productIds)
	return err
}

func productNotFoundError(id uuid.UUID) error {
	return errors.New("Product not found: " + id.String())
}
