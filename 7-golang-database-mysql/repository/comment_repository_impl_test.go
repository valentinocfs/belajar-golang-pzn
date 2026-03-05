package repository

import (
	"7-golang-database-mysql/entity"
	"context"
	"fmt"
	"testing"

	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	dsn := "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestCommentRepositoryImplInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(GetConnection())
	context := context.Background()
	comment := entity.Comment{
		Email:   "repository@gmail.com",
		Comment: "Repository Comment",
	}

	result, err := CommentRepository.Insert(context, comment)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Comment Inserted: %v\n", result)
}

func TestCommentRepositoryImplFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(GetConnection())
	context := context.Background()
	result, err := CommentRepository.FindById(context, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Comment Found: %v\n", result)
}

func TestCommentRepositoryImplFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(GetConnection())
	context := context.Background()
	result, err := CommentRepository.FindAll(context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Comments Found: %v\n", result)
}
