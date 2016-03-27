package database

import(
    "fmt"
    
    . "linq/core/database/mysql"
)

var DB MySqlDB

func UseMysql(host string, username string, password string, database string, port int) MySqlDB {
    mysqldb := MySqlDB{
        Username : username,
        Password : password,
        Database : database,
        // "user:password@tcp(localhost:5555)/dbname"
        ConnectionString : fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database),  
    }
    
    mysqldb.Ping()
    
    DB = mysqldb
    
    return mysqldb
}