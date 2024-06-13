package database

import (
	"database/sql"
	"fmt"
	"os"
)

type DB struct{}

func (db DB) Connect() (conn *sql.DB, err error) {
	conn, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/store?multiStatements=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")))

	if err != nil {
		panic(err)
	}
	return conn, err
}
