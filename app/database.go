package app

import (
	"database/sql"
	"time"

	"Mociw/golang-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/data_user?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
