package repository

type IRepository interface {
	CountAll() int
	IsExist(id string) bool
	GetAll(keyword string, length int, order int, orderDir string) []IModel
	Get(id string) IModel
	Insert(model IModel) IModel
	Update(model IModel) IModel
	Delete(model IModel) IModel
	DeleteBulk(model []string) error
}
