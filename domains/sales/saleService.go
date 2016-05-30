package sales

import (
	"errors"

	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type SaleService struct {
	repo core.IRepository
}

func NewSaleService(repo core.IRepository) SaleService {
	return SaleService{
		repo: repo,
	}
}

func (service SaleService) NewUserCart(userId uuid.UUID) (*Sale, error) {
	repo := service.repo.(saleRepository)
	return repo.NewUserCart(userId)
}

func (service SaleService) GetUserCarts(userId uuid.UUID) (*Sales, error) {
	repo := service.repo.(saleRepository)
	return repo.GetUserCarts(userId)
}

func (service SaleService) AddCartItem(sale Sale, productId uuid.UUID) error {
	repo := service.repo.(saleRepository)
	return repo.AddCartItem(sale, productId)
}

func (service SaleService) GetCartItems(sale Sale) ([]uuid.UUID, error) {
	repo := service.repo.(saleRepository)
	return repo.GetCartItems(sale)
}

func (service SaleService) CountAll() (int, error) {
	return service.repo.CountAll()
}

func (service SaleService) IsExist(id uuid.UUID) (bool, error) {
	return service.repo.IsExist(id)
}

func (service SaleService) GetAll(paging utils.Paging) (core.IModels, error) {
	return service.repo.GetAll(paging)
}

func (service SaleService) Get(id uuid.UUID) (core.IModel, error) {
	return service.repo.Get(id)
}

func (service SaleService) Create(model core.IModel) error {
	return service.repo.Insert(model)
}

func (service SaleService) Modify(model core.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		return service.repo.Update(model)
	} else {
		return saleNotFoundError()
	}
}

func (service SaleService) Remove(model core.IModel) error {
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
