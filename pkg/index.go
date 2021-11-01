package pkg

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Print(err)
	}
	firstname := session.Values["firstname"]
	block := map[string]interface{}{
		"firstname":  firstname,
		"show_block": true,
	}
	if firstname == nil {
		block["show_block"] = false
	}
	files := []string{
		dirWithHTML + "index.html",
	}
	tmp, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, block)
	if err != nil {
		log.Fatal(err)
	}
	// stoneShop := stones()
	// err = tmp.ExecuteTemplate(w, "stone", stoneShop)
	// if err != nil {
	// 	log.Print(err)
	// }
}
