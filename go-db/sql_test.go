package belajargolangdb

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

//execute mysql command
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "insert into customer(id,name) values('calvin','CALVIN')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}
	fmt.Print("succes insert into table")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "select id,name from customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id =", id)
		fmt.Println("name =", name)

	}
}

func TestAmbilData(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "select email,comment from comments"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var email, comment string

		rows.Scan(&email, &comment)
		if err != nil {
			panic(err)
		}
		fmt.Println("username= ", email)
		fmt.Println("password= ", comment)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "select id,name,email,balance,rating,created_at,birth_date,married from customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool
		rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("===========================")
		fmt.Println("id =", id)
		fmt.Println("name =", name)
		if email.Valid {
			fmt.Println("email=", email)
		}
		fmt.Println("balance =", balance)
		fmt.Println("rating =", rating)
		fmt.Println("created_at =", created_at)
		if birth_date.Valid {
			fmt.Println("email=", email)
		}
		fmt.Println("married =", married)

	}

	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"
	script := "select username from user where username ='" + username + "' and password= '" + password + "'limit 1"
	rows, err := db.QueryContext(ctx, script)
	fmt.Println(script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		fmt.Println("sukses login", username)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Print("gagal login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "calvin"
	password := "calvin"
	script := "select username from user where username =? and password= ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		fmt.Println("sukses login", username)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Print("gagal login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	username := "calvin"
	password := "calvin"
	ctx := context.Background()

	script := "insert into user(username,password) values(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	fmt.Print("succes insert into table")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "calvin@gmail.com"
	comment := "test comment"

	script := "insert into comments(email,comment) values(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Print("succes insert new comment with id ", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "insert into comments(email,comment) values(?,?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i <= 10; i++ {
		email := "calvin" + strconv.Itoa(i) + "@gmail.com"
		comment := "komen ke ->" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("comment id = ", lastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do connection
	script := "insert into comments(email,comment) values(?,?)"
	for i := 0; i <= 10; i++ {

		email := "calvin" + strconv.Itoa(i) + "@gmail.com"
		comment := "komen ke ->" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("comment id = ", id)

	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
