package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

type ChestList struct {
	ChestName    string           `json:"chestName"`
	ChestURL     string           `json:"chestURL"`
	ChestRare    string           `json:"chestRare"`
	ChestPrice   int              `json:"chestPrice"`
	ChestContent []StoneFromChest `json:"chestContent"`
}

type chestInfo struct {
	chestName    string
	chestURL     string
	chestPrice   int
	chestRare    string
	chestContent string
	chestChance  string
}

type StoneFromChest struct {
	Name        string  `json:"stoneName"`
	StoneChance float32 `json:"stoneChance,omitempty"`
	URL         string  `json:"stoneURL"`
	Rare        string  `json:"stoneRare"`
	Description string  `json:"stoneDescription,omitempty"`
}

type ChestBlock struct {
	Block     map[string]interface{}
	ChestShop []struct {
		ID    int
		Name  string
		URL   string
		Price int
		Chest string
	}
}

func NewChestBlock(block map[string]interface{}, chestShop *[]struct {
	ID    int
	Name  string
	URL   string
	Price int
	Chest string
}) *ChestBlock {
	return &ChestBlock{
		Block:     block,
		ChestShop: *chestShop}
}

func chest(db *sql.DB) *[]struct {
	ID    int
	Name  string
	URL   string
	Price int
	Chest string
} {
	rows, err := db.Query("SELECT id, name, url, price, chest FROM chests ORDER BY price DESC")
	if err != nil {
		log.Print(err)
	}
	chestShop := make([]struct {
		ID    int
		Name  string
		URL   string
		Price int
		Chest string
	}, 0, 10)
	for rows.Next() {
		var chest struct {
			ID    int
			Name  string
			URL   string
			Price int
			Chest string
		}
		err = rows.Scan(&chest.ID, &chest.Name, &chest.URL, &chest.Price, &chest.Chest)
		if err != nil {
			panic(err)
		}
		chestShop = append(chestShop, chest)
	}
	return &chestShop
}

func (h *Handler) openChest(w http.ResponseWriter, r *http.Request) {
	session, err := h.Store.Get(r, "session")
	if err != nil {
		log.Print(err)
	}
	if session.Values["userID"] == nil {
		Error := new(Error)
		Error.NewErrorMessage("???? ???????????????????????????????? ????????????????????????")
		body := new(bytes.Buffer)
		err = json.NewEncoder(body).Encode(Error)
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Fprint(w, body)
		return
	}
	if r.URL.Query().Get("id") != "" {
		var chance string
		var content string
		err := h.Storage.QueryRow("SELECT content ,chance FROM chests WHERE id=($1)", r.URL.Query().Get("id")).
			Scan(&content, &chance)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		sliceChance := strings.Split(chance, ",")
		sliceContent := strings.Split(content, ",")
		sliceFloatsLegacy := make([]float64, 0, 8)
		sliceFloatsModified := make([]float64, 0, 8)
		for _, i := range sliceChance {
			num, err := strconv.ParseFloat(i, 64)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			sliceFloatsLegacy = append(sliceFloatsLegacy, num)
		}
		sliceFloatsModified = append(sliceFloatsModified, sliceFloatsLegacy...)
		sort.Float64s(sliceFloatsModified)
		for i, j := 0, len(sliceFloatsModified)-1; i < j; i, j = i+1, j-1 {
			sliceFloatsModified[i], sliceFloatsModified[j] = sliceFloatsModified[j], sliceFloatsModified[i]
		}
		var winner int
		a := &winner
		var summary float64 = 100.0
		b := &summary
		rand.Seed(time.Now().UnixNano())
		rnd := rand.Float64() * 100
		for i, j := range sliceFloatsModified {
			if rnd <= *b && rnd >= *b-j {
				*a = i
				break
			} else if rnd <= sliceFloatsModified[len(sliceFloatsModified)-1] {
				*a = len(sliceFloatsModified) - 1
				break
			} else {
				*b = *b - j
			}
		}
		winChance := sliceFloatsModified[winner]
		var stoneWinner StoneFromChest
		for i, j := range sliceFloatsLegacy {
			if winChance == j {
				if err = h.Storage.QueryRow("SELECT name, url, rare_css, description FROM stones WHERE id=($1)", sliceContent[i]).Scan(&stoneWinner.Name, &stoneWinner.URL, &stoneWinner.Rare, &stoneWinner.Description); err != nil {
					fmt.Fprint(w, err)
				}
				buf := new(bytes.Buffer)
				if err := json.NewEncoder(buf).Encode(stoneWinner); err != nil {
					fmt.Fprint(w, err)
				}
				w.Write(buf.Bytes())
				return
			}
		}
		fmt.Fprint(w, "?????????? ?????????? ????")

	}
}

func (h *Handler) giveChests(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("id") != "" {
		var chestInfo chestInfo
		var chestList ChestList
		err := h.Storage.QueryRow("SELECT name, url, price, chest, content, chance FROM chests WHERE id=($1)", r.URL.Query().Get("id")).
			Scan(&chestInfo.chestName, &chestInfo.chestURL, &chestInfo.chestPrice, &chestInfo.chestRare, &chestInfo.chestContent, &chestInfo.chestChance)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		chestList.ChestName = chestInfo.chestName
		chestList.ChestURL = chestInfo.chestURL
		chestList.ChestRare = chestInfo.chestRare
		chestList.ChestPrice = chestInfo.chestPrice
		sliceChance := strings.Split(chestInfo.chestChance, ",")
		sliceContent := strings.Split(chestInfo.chestContent, ",")
		var wg sync.WaitGroup
		for i := range sliceContent {
			var stone StoneFromChest
			wg.Add(2)
			go func(i int) {
				b, _ := strconv.Atoi(sliceContent[i])
				err = h.Storage.QueryRow("SELECT name, url, rare_css FROM stones WHERE id=($1)", b).Scan(&stone.Name, &stone.URL, &stone.Rare)
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
		fmt.Fprint(w, "?????????? pidoras")
	}
}

func (h *Handler) chests(w http.ResponseWriter, r *http.Request) {
	session, err := h.Store.Get(r, "session")
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
	ChestBlock := NewChestBlock(block, chest(h.Storage))
	tmp, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, *ChestBlock)
	if err != nil {
		fmt.Fprint(w, err)
	}
}
