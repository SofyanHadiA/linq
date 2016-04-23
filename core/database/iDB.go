package database

import (
	"database/sql"
	"linq/core/repository"
	
    "github.com/jmoiron/sqlx"
)

type IDB interface {
	Ping() bool
	Select(query string, model repository.IModel) error
	ResolveSingle(query string, args ...interface{}) *sqlx.Row
	Resolve(query string, args ...interface{}) *sqlx.Rows
	Execute(query string, model repository.IModel) sql.Result
}
