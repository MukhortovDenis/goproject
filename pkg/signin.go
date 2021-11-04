package pkg

import (
	"fmt"
	"net/http"
	"text/template"
)

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(dirWithHTML + "signin.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, err)
	}
}

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(dirWithHTML + "signup.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, err)
	}
}
