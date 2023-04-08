package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"pionirx/config"
	"pionirx/helper"
)

var db *sql.DB

func Init() {
	cfg := config.Init()
	connectionString := cfg.DbUsername + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":" + cfg.DbPort + ")/" + cfg.DbName
	db, err := sql.Open("mysql", connectionString)

	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)
}

func CreateDb() *sql.DB {
	return db
}
