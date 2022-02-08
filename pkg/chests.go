package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/template"
)

type ChestList struct {
	ChestName    string           `json:"chestName"`
	ChestURL     string           `json:"chestURL"`
	ChestPrice   int           `json:"chestPrice"`
	ChestContent []StoneFromChest `json:"chestContent"`
}

type chestInfo struct {
	chestName    string
	chestURL     string
	chestPrice   int
	chestContent string
	chestChance  string
}

type StoneFromChest struct {
	Name        string  `json:"stoneName"`
	StoneChance float32 `json:"stoneChance"`
	URL         string  `json:"stoneURL"`
	Rare        string  `json:"stoneRare"`
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
	if r.URL.Query().Get("id") != "" {
		db, err := sql.Open("postgres", dbConn)
		if err != nil {
			fmt.Fprint(w, err)
		}
		defer db.Close()
		var chestInfo chestInfo
		var chestList ChestList
		err = db.QueryRow("SELECT name, url, price, content, chance FROM chests WHERE id=($1)", r.URL.Query().Get("id")).
			Scan(&chestInfo.chestName, &chestInfo.chestURL, &chestInfo.chestPrice, &chestInfo.chestContent, &chestInfo.chestChance)
		if err != nil {
			fmt.Fprint(w, err)
		}
		chestList.ChestName = chestInfo.chestName
		chestList.ChestURL = chestInfo.chestURL
		chestList.ChestPrice = chestInfo.chestPrice
		sliceChance := strings.Split(chestInfo.chestChance, ",")
		sliceContent := strings.Split(chestInfo.chestContent, ",")
		var wg sync.WaitGroup
		for i := range sliceContent {
			var stone StoneFromChest
			wg.Add(2)
			go func(i int) {
				b, _ := strconv.Atoi(sliceContent[i])
				err = db.QueryRow("SELECT name, url, rare_css FROM stones WHERE id=($1)", b).Scan(&stone.Name, &stone.URL, &stone.Rare)
				if err != nil {
					fmt.Fprint(w, err)
				}
				wg.Done()
			}(i)
			go func(i int) {
				b, _ := strconv.ParseFloat(sliceChance[i], 32)
				c := float32(b)
				stone.StoneChance = c
				wg.Done()
			}(i)
			wg.Wait()
			chestList.ChestContent = append(chestList.ChestContent, stone)
		}
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(chestList); err != nil {
			fmt.Fprint(w, err)
		}
		w.Write(buf.Bytes())
	} else {
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
