package database

import (
	"database/sql"
	"bitbucket.org/sofyan_a/linq.im/core/repository"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type IDB interface {
	Ping() (bool, error)
	ResolveSingle(query string, args ...interface{}) (*sqlx.Row, error)
	Resolve(query string, args ...interface{}) (*sqlx.Rows, error)
	Execute(query string, model repository.IModel) (*sql.Result, error)
	ExecuteArgs(query string, params ...interface{}) (*sql.Result, error)
	ExecuteBulk(query string, data []uuid.UUID) (*sql.Result, error)
}
