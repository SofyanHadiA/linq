package database

import (
	"database/sql"
	"fmt"

	"linq/core/utils"
	"linq/core/repository"

	_ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

type mySqlDB struct {
	Host             string
	Port             int
	Username         string
	Password         string
	Database         string
	ConnectionString string
}

func MySqlDB(host string, username string, password string, database string, port int) IDB {
	DB := mySqlDB{
		Username: username,
		Password: password,
		Database: database,
		ConnectionString: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database),
	}
	
	go DB.Ping()
	
	return DB
}

func (mysql mySqlDB) Ping() bool {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	err = db.Ping()
	if err != nil {
		utils.Log.Fatal(err.Error(), mysql.ConnectionString)
	} else {
		utils.Log.Info("Connected to mysql server", mysql.ConnectionString)
	}

	return true
}

func (mysql mySqlDB) Select(query string, model repository.IModel) error{
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)

	defer db.Close()

    err = db.Select(&model, query)
    return err
}

func (mysql mySqlDB) ResolveSingle(query string, args ...interface{}) *sqlx.Row {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)

	defer db.Close()
	
	row := db.QueryRowx(query, args...)

	return row
}


func (mysql mySqlDB) Resolve(query string, args ...interface{}) *sqlx.Rows {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)
	defer db.Close()

	rows, err := db.Queryx(query, args...)
	utils.HandleWarn(err)

	return rows
}

func (mysql mySqlDB) Execute(query string, model repository.IModel) sql.Result{
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)
	defer db.Close()
	
	result, err := db.NamedExec(query, model)
	utils.HandleWarn(err)
	
	return result
}
