package belajargolangdb

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:748159263@tcp(localhost:3306)/belajar_golang_db")
	if err != nil {
		panic(err)
	}
	//gunakan db

	defer db.Close()
}
