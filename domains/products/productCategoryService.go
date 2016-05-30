package products

import (
	"errors"

	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type ProductCategoryService struct {
	repo core.IRepository
}

func NewProductCategoryService(repo core.IRepository) ProductCategoryService {
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

func (service ProductCategoryService) GetAll(paging utils.Paging) (core.IModels, error) {
	return service.repo.GetAll(paging)
}

func (service ProductCategoryService) Get(id uuid.UUID) (core.IModel, error) {
	return service.repo.Get(id)
}

func (service ProductCategoryService) Create(model core.IModel) error {
	return service.repo.Insert(model)
}

func (service ProductCategoryService) Modify(model core.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		return service.repo.Update(model)
	} else {
		return productCategoryNotFoundError()
	}
}

func (service ProductCategoryService) Remove(model core.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		err := service.repo.Delete(model)
		return err
	} else {
		return productCategoryNotFoundError()
	}
}

func (service ProductCategoryService) RemoveBulk(productCategoryIds []uuid.UUID) error {
	for _, id := range productCategoryIds {
		if exist, _ := service.repo.IsExist(id); !exist {
			return productCategoryNotFoundError()
		}
	}

	err := service.repo.DeleteBulk(productCategoryIds)
	return err
}

func productCategoryNotFoundError() error {
	return errors.New("ProductCategoryNotFound")
}
