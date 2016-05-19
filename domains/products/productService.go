package products

import (
	"errors"
	"fmt"

	"github.com/SofyanHadiA/linq/core/repository"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type ProductService struct {
	repo repository.IRepository
}

func NewProductService(repo repository.IRepository) ProductService {
	return ProductService{
		repo: repo,
	}
}

func (service ProductService) CountAll() (int, error) {
	return service.repo.CountAll()
}

func (service ProductService) IsExist(id uuid.UUID) (bool, error) {
	return service.repo.IsExist(id)
}

func (service ProductService) GetAll(paging utils.Paging) (repository.IModels, error) {
	return service.repo.GetAll(paging)
}

func (service ProductService) Get(id uuid.UUID) (repository.IModel, error) {
	return service.repo.Get(id)
}

func (service ProductService) Create(model repository.IModel) error {
	return service.repo.Insert(model)
}

func (service ProductService) Modify(model repository.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		return service.repo.Update(model)
	} else {
		return productNotFoundError()
	}
}

func (service ProductService) UpdateProductPhoto(model repository.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		productRepo := service.repo.(productRepository)

		err := productRepo.UpdateProductPhoto(model)

		return err
	} else {
		return productNotFoundError()
	}
}

func (service ProductService) Remove(model repository.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		err := service.repo.Delete(model)

		return err
	} else {
		return productNotFoundError()
	}
}

func (service ProductService) RemoveBulk(productIds []uuid.UUID) error {
	for _, id := range productIds {
		if exist, _ := service.repo.IsExist(id); exist {
			return productNotFoundError()
		}
	}

	err := service.repo.DeleteBulk(productIds)
	return err
}

func productNotFoundError() error {
	return errors.New("ProductNotFound")
}

func sha1ToString(c [20]byte) string {
	return string(fmt.Sprintf("%x", c))
}
