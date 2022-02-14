package pkg

import (
	"database/sql"
	"log"
)

func stones(db *sql.DB) *[]struct {
	ID          int
	Name        string
	URL         string
	Description string
	Price       int
	RareCss     string
	Rare        string
} {
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
	}, 0, 50)
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
	return &stoneStore
}
