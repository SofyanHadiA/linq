package database

import (
	"database/sql"
	"fmt"

	"linq/core/utils"

	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", mysql.ConnectionString)
	err = db.Ping()
	if err != nil {
		utils.Log.Fatal(err.Error(), mysql.ConnectionString)
	} else {
		utils.Log.Info("Connected to mysql server", mysql.ConnectionString)
	}

	return true
}

func (mysql mySqlDB) ResolveSingle(query string, args ...interface{}) *sql.Row {
	db, err := sql.Open("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)

	defer db.Close()

	var row = &sql.Row{}

	if len(args) > 0 {
		row = db.QueryRow(query, args...)
	} else {
		row = db.QueryRow(query)
	}

	return row
}


func (mysql mySqlDB) Resolve(query string, args ...interface{}) *sql.Rows {
	db, err := sql.Open("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)
	defer db.Close()

	var rows = &sql.Rows{}

	if len(args) > 0 {
		rows, err = db.Query(query, args...)
	} else {
		rows, err = db.Query(query)
	}

	utils.HandleWarn(err)

	return rows
}

func (mysql mySqlDB) Execute(query string, args ...interface{}) sql.Result {
	db, err := sql.Open("mysql", mysql.ConnectionString)
	utils.HandleWarn(err)

	defer db.Close()

	stmtOut, err := db.Prepare(query)
	
	utils.HandleWarn(err)
	defer stmtOut.Close()
	
	res, err := stmtOut.Exec(args...)
	utils.HandleWarn(err)

	return res
}
