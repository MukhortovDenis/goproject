package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func stones() []struct {
	ID          int
	Name        string
	URL         string
	Description string
	Price       int
	RareCss     string
	Rare        string
} {
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	rows, err := db.Query("SELECT * FROM stones ORDER BY price DESC")
	if err != nil {
		log.Print(err)
	}
	stoneStore := make([]struct {
		ID          int
		Name        string
		URL         string
		Description string
		Price       int
		RareCss     string
		Rare        string
	}, 0, 10)
	for rows.Next() {
		var stone struct {
			ID          int
			Name        string
			URL         string
			Description string
			Price       int
			RareCss     string
			Rare        string
		}
		err = rows.Scan(&stone.ID, &stone.Name, &stone.URL, &stone.Description, &stone.Price, &stone.RareCss, &stone.Rare)
		if err != nil {
			panic(err)
		}
		stoneStore = append(stoneStore, stone)

	}
	return stoneStore
}
