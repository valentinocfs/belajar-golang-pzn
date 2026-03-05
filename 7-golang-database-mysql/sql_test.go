package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('1', 'John Doe')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data")
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id, "|", "Name:", name)
	}
	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var married bool
		var createdAt time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		// handle null value
		var emailValue string
		if email.Valid {
			emailValue = email.String
		}

		// handle null value
		var birthDateValue time.Time
		if birthDate.Valid {
			birthDateValue = birthDate.Time
		}

		fmt.Println("ID:", id, "|", "Name:", name, "|", "Email:", emailValue, "|", "Balance:", balance, "|", "Rating:", rating, "|", "Birth Date:", birthDateValue, "|", "Married:", married, "|", "Created At:", createdAt)
	}
	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// implementasi sql injection
	username := "admin'; #"
	password := "salah"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success, Welcome", username)
	} else {
		fmt.Println("Login Failed")
	}

	defer rows.Close()
}

func TestSqlWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// implementasi sql injection
	username := "admin'; #"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success, Welcome", username)
	} else {
		fmt.Println("Login Failed")
	}

	defer rows.Close()
}

func TestInsertQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "john.doe@example.com"
	comment := "Awesome"

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data with ID:", lastInsertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	stmt, err := db.PrepareContext(ctx, "INSERT INTO comments(email, comment) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 1; i <= 10; i++ {
		email := "john.doe@example.com"
		comment := fmt.Sprintf("Comment %d", i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Insert Data with ID:", lastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	for i := 1; i <= 10; i++ {
		email := "john.doe@example.com"
		comment := fmt.Sprintf("Comment %d", i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Insert Comment with ID:", lastInsertId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
