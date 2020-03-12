package database

import(
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
const(
    driverName string = "mysql"
    users string = "admin"
    password string = "password"
    host string = "127.0.0.1:3306"
    databaseName string = "mydb"
    dataSouceName string = users+":"+password+"@tcp("+host+")/"+databaseName
)

type MySQLCli struct {
  db *sql.DB
}

var instanceMySQLCli *MySQLCli = nil

func Connect() (db *sql.DB, err error) {
    if instanceMySQLCli == nil {
        instanceMySQLCli = new(MySQLCli)
        var err error

        instanceMySQLCli.db, err = sql.Open(driverName, dataSouceName)
        
        if err != nil {
           return nil, err
        }
    }

    return instanceMySQLCli.db, nil
}

func Close() {
    if instanceMySQLCli != nil {
       instanceMySQLCli.db.Close()

    }
}