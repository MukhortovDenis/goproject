package pkg

import (
	"fmt"
	"net/http"
	"text/template"
)

func (h *Handler) settings(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(dirWithHTML + "settings.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, err)
	}
}

func (h *Handler) settingsAppearance(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(dirWithHTML + "settings-appearance.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, err)
	}
}
