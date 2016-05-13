package services

import (
	. "bitbucket.org/sofyan_a/linq.im/core/repository"
	"bitbucket.org/sofyan_a/linq.im/core/utils"

	"github.com/satori/go.uuid"
)

type IService interface {
	CountAll() (int, error)
	IsExist(id uuid.UUID) (bool, error)
	GetAll(paging utils.Paging) (IModels, error)
	Get(id uuid.UUID) (IModel, error)
	Create(model IModel) error
	Modify(model IModel) error
	Remove(model IModel) error
	RemoveBulk(model []uuid.UUID) error
}
