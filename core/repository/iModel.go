package repository

type IModel interface {
	GetId() int
	InsertVal() []interface{}
	UpdateVal() []interface{}
}