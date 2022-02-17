package pkg

import (
	"fmt"
	"net/http"
	"text/template"
)

func (h *Handler) inventory(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(dirWithHTML + "inventory.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, err)
	}
}
