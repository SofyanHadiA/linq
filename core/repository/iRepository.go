package repository

type IModel interface{
    GetId() int
}

type IRepository interface {
    CountAll() int
    GetAll() []IModel
    Get(id int) IModel
}
