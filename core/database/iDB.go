package database

import (
	"database/sql"
)

type IDB interface {
	Ping() bool
	ResolveSingle(query string, args ...interface{}) *sql.Row
	Resolve(query string, args ...interface{}) *sql.Rows
	Execute(query string, args ...interface{}) sql.Result
}
