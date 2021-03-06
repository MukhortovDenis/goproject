package pkg

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type IndexBlock struct {
	Block     map[string]interface{}
	StoneShop []struct {
		ID          int
		Name        string
		URL         string
		Description string
		Price       int
		RareCss     string
		Rare        string
	}
}

func NewIndexBlock(block map[string]interface{}, stoneShop *[]struct {
	ID          int
	Name        string
	URL         string
	Description string
	Price       int
	RareCss     string
	Rare        string
}) *IndexBlock {
	return &IndexBlock{
		Block:     block,
		StoneShop: *stoneShop,
	}
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
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
		dirWithHTML + "index.html",
		dirWithHTML + "stone-temp.html",
	}
	indexBlock := NewIndexBlock(block, stones(h.Storage))
	tmp, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, *indexBlock)
	if err != nil {
		log.Fatal(err)
	}
}
