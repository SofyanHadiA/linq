package mysql

import(
    "database/sql"

    utils "linq/core/utils"
    
    _ "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
    Host string
    Port int
    Username string
    Password string
    Database string
    ConnectionString string
    query string
}

func (mysql MySqlDB) Ping() bool{
    db, err := sql.Open("mysql", mysql.ConnectionString) 
    err = db.Ping()
    if err != nil {
        utils.Log.Fatal(err.Error(), mysql.ConnectionString)
    }else{
        utils.Log.Info("Connected to mysql server", mysql.ConnectionString)
    }
    
    return true
}

func (mysql MySqlDB) Resolve(query string) *sql.Rows{
    db, err := sql.Open("mysql", mysql.ConnectionString) 
    defer db.Close()
    
    rows, err := db.Query(query)
    utils.HandleWarn(err)
    
    return rows
}

func (mysql MySqlDB) Execute(query string) sql.Result{
    db, err := sql.Open("mysql", mysql.ConnectionString) 

    stmtOut, err := db.Prepare(query)
    utils.HandleWarn(err)
    defer stmtOut.Close()
    
    res, err := stmtOut.Exec()
    utils.HandleWarn(err)
    
    return res
}
