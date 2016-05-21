package repository

import (
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type IRepository interface {
	CountAll() (int, error)
	IsExist(id uuid.UUID) (bool, error)
	GetAll(paging utils.Paging) (IModels, error)
	Get(id uuid.UUID) (IModel, error)
	Insert(model IModel) error
	Update(model IModel) error
	Delete(model IModel) error
	DeleteBulk(model []uuid.UUID) error
}
