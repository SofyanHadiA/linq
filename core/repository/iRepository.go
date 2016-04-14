package repository

type IRepository interface {
	CountAll() int
	GetAll() []IModel
	Get(id int) IModel
	Insert(model IModel) IModel
	Update(model IModel) IModel
}
