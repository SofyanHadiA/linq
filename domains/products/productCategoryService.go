package products

import (
	"errors"
	"fmt"

	"github.com/SofyanHadiA/linq/core/repository"
	"github.com/SofyanHadiA/linq/core/services"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type ProductCategoryService struct {
	repo          repository.IRepository
	uploadService services.IUploadService
}

func NewProductCategoryService(repo repository.IRepository, uploadService services.IUploadService) ProductCategoryService {
	return ProductCategoryService{
		repo:          repo,
		uploadService: uploadService,
	}
}

func (service ProductCategoryService) CountAll() (int, error) {
	return service.repo.CountAll()
}

func (service ProductCategoryService) IsExist(id uuid.UUID) (bool, error) {
	return service.repo.IsExist(id)
}

func (service ProductCategoryService) GetAll(paging utils.Paging) (repository.IModels, error) {
	return service.repo.GetAll(paging)
}

func (service ProductCategoryService) Get(id uuid.UUID) (repository.IModel, error) {
	return service.repo.Get(id)
}

func (service ProductCategoryService) Create(model repository.IModel) error {
	return service.repo.Insert(model)
}

func (service ProductCategoryService) Modify(model repository.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		return service.repo.Update(model)
	} else {
		return productCategoryNotFoundError()
	}
}

func (service ProductCategoryService) UpdateProductCategoryPhoto(model repository.IModel, imageString string) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		fileName := model.GetId().String() + ".png"
		err := service.uploadService.UploadImage(imageString, fileName)

		if err == nil {
			productCategory, _ := model.(*ProductCategory)

			productCategory.Image.String = fileName
			productCategory.Image.Valid = true

			productCategoryRepo := service.repo.(productCategoryRepository)
			err = productCategoryRepo.UpdateProductCategoryPhoto(productCategory)
		}

		return err
	} else {
		return productCategoryNotFoundError()
	}
}

func (service ProductCategoryService) Remove(model repository.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		err := service.repo.Delete(model)

		return err
	} else {
		return productCategoryNotFoundError()
	}
}

func (service ProductCategoryService) RemoveBulk(productCategoryIds []uuid.UUID) error {
	for _, id := range productCategoryIds {
		if exist, _ := service.repo.IsExist(id); exist {
			return productCategoryNotFoundError()
		}
	}

	err := service.repo.DeleteBulk(productCategoryIds)
	return err
}

func productCategoryNotFoundError() error {
	return errors.New("ProductCategoryNotFound")
}

func sha1ToString(c [20]byte) string {
	return string(fmt.Sprintf("%x", c))
}
