package repository

import (
	"reflect"

	"github.com/satori/go.uuid"
)

type IModel interface {
	GetId() uuid.UUID
}

type IModels interface {
	GetLength() int
}

func GetDBField(model IModel) map[string]interface{} {
	val := reflect.ValueOf(model).Elem()

	dbField := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		if tag.Get("db") != "" {
			dbField[tag.Get("db")] = val.Field(i)
		}
	}

	return dbField
}
