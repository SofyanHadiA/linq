package products

import (
	"errors"

	"github.com/SofyanHadiA/linq/core/repository"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type ProductCategoryService struct {
	repo repository.IRepository
}

func NewProductCategoryService(repo repository.IRepository) ProductCategoryService {
	return ProductCategoryService{
		repo: repo,
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
