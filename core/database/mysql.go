package database

import (
	"database/sql"
	"errors"
	"fmt"

	"linq/core/repository"
	"linq/core/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
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
		Username:         username,
		Password:         password,
		Database:         database,
		ConnectionString: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database),
	}

	_, err := DB.Ping()
	if err != nil {
		utils.Log.Fatal(err.Error(), DB.ConnectionString)
	}

	return DB
}

func (mysql mySqlDB) Ping() (bool, error) {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	if err == nil {
		if err = db.Ping(); err != nil {
			utils.HandleWarn(err)
			return false, err
		} else {
			utils.Log.Info("Connected to mysql server", mysql.ConnectionString)
			return true, nil
		}
	} else {
		utils.HandleWarn(err)
		return false, err
	}
}

func (mysql mySqlDB) ResolveSingle(query string, args ...interface{}) (*sqlx.Row, error) {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)
	defer db.Close()

	if err == nil {
		row := db.QueryRowx(query, args...)
		utils.HandleWarn(err)

		return row, err
	} else {
		return nil, dbConectError()
	}
}

func (mysql mySqlDB) Resolve(query string, args ...interface{}) (*sqlx.Rows, error) {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)
	defer db.Close()

	if err == nil {
		rows, err := db.Queryx(query, args...)
		utils.HandleWarn(err)

		return rows, err
	} else {
		return nil, dbConectError()
	}
}

func (mysql mySqlDB) Execute(query string, model repository.IModel) (*sql.Result, error) {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)
	defer db.Close()

	if err == nil {
		result, err := db.NamedExec(query, model)
		utils.HandleWarn(err)
		return &result, err
	} else {
		return nil, dbConectError()
	}
}

func (mysql mySqlDB) ExecuteArgs(query string, params ...interface{}) (*sql.Result, error) {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)
	defer db.Close()

	if err == nil {
		result, err := db.Exec(query, params...)
		utils.HandleWarn(err)
		return &result, err
	} else {
		return nil, dbConectError()
	}
}

func (mysql mySqlDB) ExecuteBulk(query string, data []uuid.UUID) (*sql.Result, error) {
	db, err := sqlx.Connect("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)
	defer db.Close()

	if err == nil {
		query, args, err := sqlx.In(query, data)
		utils.HandleWarn(err)
		query = db.Rebind(query)
		result := db.MustExec(query, args...)
		return &result, err
	} else {
		return nil, dbConectError()
	}
}

func dbConectError() error {
	return errors.New("CannotConnectToDatabase")
}
