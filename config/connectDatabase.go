package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/goschool")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success connect to database")

	return db
}
