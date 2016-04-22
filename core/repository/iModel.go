package repository

type IModel interface {
	GetId() string
	InsertVal() []interface{}
	UpdateVal() []interface{}
}