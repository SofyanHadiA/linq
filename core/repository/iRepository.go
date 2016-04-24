package repository

import (
	"database/sql"

	"github.com/satori/go.uuid"
)

type IRepository interface {
	CountAll() int
	IsExist(id uuid.UUID) bool
	GetAll(keyword string, length int, order int, orderDir string) IModels
	Get(id uuid.UUID) IModel
	Insert(model IModel) IModel
	Update(model IModel) IModel
	Delete(model IModel) IModel
	DeleteBulk(model []uuid.UUID) sql.Result
}

func ResolveSelectFields(model IModel) string {
	fields := GetDBField(model)

	var result string

	i := 0

	for key := range fields {
		if i > 0 {
			result += ", "
		}
		result += "`" + key + "`"
		i++
	}

	return result
}
