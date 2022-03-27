package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var MainDB *sqlx.DB

func ConnectDB() {
	MainDB := sqlx.MustConnect("mysql", "")
	MainDB.SetMaxOpenConns(500)
	MainDB.SetMaxIdleConns(100)
}
