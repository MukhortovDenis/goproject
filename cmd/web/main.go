package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "mysql:123@tcp(127.0.0.1:3306)/stoneshop")
	if err != nil {
		panic(err)
	}
	insert, err := db.Query("INSERT INTO `users` (`login`, `password`) VALUES('ZhizhaDon', 'zhizhadon282829')")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	defer db.Close()

	mainHandle()
	fmt.Println("я не лох")

}
