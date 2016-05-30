package core

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type IDB interface {
	Ping() (bool, error)
	ResolveSingle(query string, args ...interface{}) (*sqlx.Row, error)
	Resolve(query string, args ...interface{}) (*sqlx.Rows, error)
	Execute(query string, model IModel) (*sql.Result, error)
	ExecuteArgs(query string, params ...interface{}) (*sql.Result, error)
	ExecuteBulk(query string, data []uuid.UUID) (*sql.Result, error)
}
