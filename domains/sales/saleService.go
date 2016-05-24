package sales

import (
	"errors"

	"github.com/SofyanHadiA/linq/core/repository"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type SaleService struct {
	repo repository.IRepository
}

func NewSaleService(repo repository.IRepository) SaleService {
	return SaleService{
		repo: repo,
	}
}

func (service SaleService) CountAll() (int, error) {
	return service.repo.CountAll()
}

func (service SaleService) IsExist(id uuid.UUID) (bool, error) {
	return service.repo.IsExist(id)
}

func (service SaleService) GetAll(paging utils.Paging) (repository.IModels, error) {
	return service.repo.GetAll(paging)
}

func (service SaleService) Get(id uuid.UUID) (repository.IModel, error) {
	return service.repo.Get(id)
}

func (service SaleService) Create(model repository.IModel) error {
	return service.repo.Insert(model)
}

func (service SaleService) Modify(model repository.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		return service.repo.Update(model)
	} else {
		return saleNotFoundError()
	}
}

func (service SaleService) Remove(model repository.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		err := service.repo.Delete(model)

		return err
	} else {
		return saleNotFoundError()
	}
}

func (service SaleService) RemoveBulk(saleIds []uuid.UUID) error {
	for _, id := range saleIds {
		if exist, _ := service.repo.IsExist(id); exist {
			return saleNotFoundError()
		}
	}

	err := service.repo.DeleteBulk(saleIds)
	return err
}

func saleNotFoundError() error {
	return errors.New("SaleNotFound")
}
