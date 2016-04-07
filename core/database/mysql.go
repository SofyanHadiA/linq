package database

import(
    "fmt"
    "database/sql"

    "linq/core/utils"
    
    _ "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
    Host string
    Port int
    Username string
    Password string
    Database string
    ConnectionString string
}


func NewMysqlDB(host string, username string, password string, database string, port int) IDB {
    DB := MySqlDB{
        Username : username,
        Password : password,
        Database : database,
        // Format : "user:password@tcp(localhost:5555)/dbname"
        ConnectionString : fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database),  
    }
    DB.Ping()
    return DB
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

func (mysql MySqlDB) Resolve(query string, args ...interface{}) *sql.Rows{
    db, err := sql.Open("mysql", mysql.ConnectionString) 
    defer db.Close()
    
    var rows = &sql.Rows{}
    
    if len(args) > 0 {
        rows, err = db.Query(query, args...)
    }else{
        rows, err = db.Query(query)
    }
    
    utils.HandleWarn(err)
    
    return rows
}

func (mysql MySqlDB) Execute(query string) sql.Result{
    db, err := sql.Open("mysql", mysql.ConnectionString) 
    defer db.Close()

    stmtOut, err := db.Prepare(query)
    utils.HandleWarn(err)
    defer stmtOut.Close()
    
    res, err := stmtOut.Exec()
    utils.HandleWarn(err)
    
    return res
}
