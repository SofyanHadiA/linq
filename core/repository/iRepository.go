package repository

type IRepository interface {
	CountAll() int
	IsExist(id int) bool
	GetAll(keyword string, order string, orderDir string) []IModel
	Get(id int) IModel
	Insert(model IModel) IModel
	Update(model IModel) IModel
	Delete(model IModel) IModel
}
