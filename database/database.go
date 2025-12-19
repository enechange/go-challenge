package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	config "go-challenge/config"
)

var (
	dbMy *sql.DB
	err  error
)

func Init() {
	openMySQL()
}

func Close() {
	if dbMy != nil {
		dbMy.Close()
	}
}

func openMySQL() {
	cnf := config.GetConfig()

	database := cnf.MySQLDatabase

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cnf.MySQLUser,
		cnf.MySQLPassword,
		cnf.MySQLHost,
		cnf.MySQLPort,
		database,
	)

	dbMy, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	if err = dbMy.Ping(); err != nil {
		panic(err)
	}
}
