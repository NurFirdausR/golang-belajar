package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConn(t *testing.T) {
	db, err := sql.Open("mysql", "root:Bismillah@123@tcp(localhost:3306)/golang_pos?parseTime=true")

	if err != nil {
		panic(err)
	}
	defer db.Close()

}
