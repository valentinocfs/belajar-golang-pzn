package app

import (
	"database/sql"
	"golang-database-migration/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang_database_migration?parseTime=True")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/belajar_golang_database_migration" -path db/migrations up
// migrate -database "mysql://root:@tcp(127.0.0.1:3306)/belajar_golang_database_migration" -path db/migrations down
