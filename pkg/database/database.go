package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SqlDatabase struct {
	Conn *sqlx.DB
}

func MySqlInitialize(host, username, password, database, port string) (SqlDatabase, error) {
	db := SqlDatabase{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()

	if err != nil {
		return db, err
	}

	return db, nil
}
