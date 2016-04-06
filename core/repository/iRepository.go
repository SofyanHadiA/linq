package repository

type IModel interface{
    SetupModel() IModel
}

type IRepository interface {
    CountAll() int
    GetAll() []IModel
}
