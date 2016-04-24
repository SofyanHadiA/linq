package database

import (
	"database/sql"
	"linq/core/repository"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type IDB interface {
	Ping() bool
	Select(query string, model repository.IModel) error
	ResolveSingle(query string, args ...interface{}) *sqlx.Row
	Resolve(query string, args ...interface{}) *sqlx.Rows
	Execute(query string, model repository.IModel) sql.Result
	ExecuteBulk(query string, data []uuid.UUID) sql.Result
}
