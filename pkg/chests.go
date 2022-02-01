package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type ChestList struct {
	ChestName    string           `json:"treasureName"`
	ChestContent []StoneFromChest `json:"treasureContent"`
}

type StoneFromChest struct {
	Name string `json:"stoneName"`
	ID   uint   `json:"stoneID"`
	URL  string `json:"stoneURL"`
	Rare string `json:"stoneRare"`
}

type ChestBlock struct {
	Block     map[string]interface{}
	ChestShop []struct {
		ID    int
		Name  string
		URL   string
		Price int
		Chest string
		Rare  string
	}
}

func NewChestBlock(block map[string]interface{}, chestShop *[]struct {
	ID    int
	Name  string
	URL   string
	Price int
	Chest string
	Rare  string
}) *ChestBlock {
	return &ChestBlock{
		Block:     block,
		ChestShop: *chestShop}
}

func chest() *[]struct {
	ID    int
	Name  string
	URL   string
	Price int
	Chest string
	Rare  string
} {
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM chests ORDER BY price DESC")
	if err != nil {
		log.Print(err)
	}
	chestShop := make([]struct {
		ID    int
		Name  string
		URL   string
		Price int
		Chest string
		Rare  string
	}, 0, 10)
	for rows.Next() {
		var chest struct {
			ID    int
			Name  string
			URL   string
			Price int
			Chest string
			Rare  string
		}
		err = rows.Scan(&chest.ID, &chest.Name, &chest.URL, &chest.Price, &chest.Chest, &chest.Rare)
		if err != nil {
			panic(err)
		}
		chestShop = append(chestShop, chest)
	}
	return &chestShop
}

func (h *Handler) giveChests(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("id") == "1" {
		chest := ChestList{ChestName: "The Chest of the Legendary Quarry",
			ChestContent: []StoneFromChest{
				{Name: "Хуета", ID: 1, URL: "static/images/common/chest.png", Rare: "poop"},
				{Name: "Залупня", ID: 2, URL: "static/images/common/chest.png", Rare: "poop"}}}

		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(chest); err != nil {
			fmt.Fprint(w, err)
		}
		w.Write(buf.Bytes())
	}else{
		fmt.Fprint(w, "автор pidoras")
	}
}

func (h *Handler) chests(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Print(err)
	}
	firstname := session.Values["firstname"]
	block := map[string]interface{}{
		"Firstname":  firstname,
		"Show_block": true,
	}
	if firstname == nil {
		block["Show_block"] = false
	}
	files := []string{
		dirWithHTML + "chests.html",
		dirWithHTML + "chest-temp.html",
	}
	ChestBlock := NewChestBlock(block, chest())
	tmp, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, *ChestBlock)
	if err != nil {
		fmt.Fprint(w, err)
	}
}
