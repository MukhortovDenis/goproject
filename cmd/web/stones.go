package main

import (
	"database/sql"
	"fmt"
	"goproject/pkg"
	"log"
	"os"
)

func stones() map[string]interface{} {
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	rows, err := db.Query("SELECT * FROM stones")
	if err != nil {
		log.Print(err)
	}
	stoneStore := make(map[string]interface{})
	for rows.Next() {
		stone := pkg.Stone{}
		err = rows.Scan(&stone.ID, &stone.Name, &stone.URL, &stone.Description, &stone.Price, &stone.Rare)
		if err != nil {
			panic(err)
		}
		name := stone.Name
		stoneStore[name] = stone
	}
	return stoneStore
}
