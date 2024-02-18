package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConn()

	defer db.Close()

	ctx := context.Background()
	id := 2
	username := "qwen"
	password := "password"
	script := "INSERT INTO user(id, username,password) VALUES (?, ?,?)"
	// script := "INSERT INTO customer(id, name,email,balance,rating,birth_date,married) VALUES ('nur', 'Nur','nurfirdaus.5000@gmail.com',10000,5,'2000-01-04',false)"

	_, err := db.ExecContext(ctx, script, id, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConn()

	defer db.Close()
	ctx := context.Background()
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		// var created_at string
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)

		}
		fmt.Println("id", id)
		fmt.Println("name", name)
		if email.Valid {
			fmt.Println("email", email.String)
		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		if birth_date.Valid {
			fmt.Println("birthdate", birth_date.Time)
		}
		fmt.Println("created_at", created_at)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConn()
	defer db.Close()

	ctx := context.Background()

	username := "qwen"
	password := "password"

	script := "SELECT username from user where username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Berhasil login", username)

	} else {
		fmt.Println("Gagal login")
	}

}
func TestAutoIncrement(t *testing.T) {
	db := GetConn()

	defer db.Close()

	ctx := context.Background()

	email := "nurfirdaus.5000@gmail.com"
	comment := "woy sadar"

	script := "INSERT INTO comments(email,comment) values(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("id last ", insertedId)
}

func TestPreparestatement(t *testing.T) {
	db := GetConn()

	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email,comment) values(?,?)"

	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "nurfirdaus.5000" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("id last", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConn()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transanction

	script := "INSERT INTO comments(email,comment) values(?,?)"
	for i := 0; i < 10; i++ {
		email := "nurfirdaus.5000" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment" + strconv.Itoa(i) + " transaction"

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("id latest", id)

	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
