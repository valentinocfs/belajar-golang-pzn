package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	fmt.Println("Initialize Driver Database")
}

func TestOpenConnection(t *testing.T) {
	dsn := "root:@tcp(localhost:3306)/belajar_golang_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Success Open Connection to Database")
}
