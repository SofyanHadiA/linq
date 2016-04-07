package database

import(
    "database/sql"
)

type IDB interface{
    Ping() bool
    Resolve(query string, params ...interface{}) *sql.Rows
    Execute(query string) sql.Result
}