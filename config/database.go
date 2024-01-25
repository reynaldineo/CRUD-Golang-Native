package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBCOnnection() (*sql.DB, error)  {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "crud_go_native"

	db, err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}