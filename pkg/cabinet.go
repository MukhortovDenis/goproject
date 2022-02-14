package pkg

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (h *Handler) cabinet(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(dirWithHTML + "cabinet.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) cabinetInfo(w http.ResponseWriter, r *http.Request) {
	session, err := h.Store.Get(r, "session")
	if err != nil {
		log.Print(err)
	}
	firstname := session.Values["firstname"]
	email := session.Values["email"]
	userInfo := map[string]interface{}{
		"firstname": firstname,
		"email":     email,
	}
	tmp, err := template.ParseFiles(dirWithHTML + "cabinet-info.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, userInfo)
	if err != nil {
		fmt.Fprint(w, err)
	}
}

func (h *Handler) cabinetPassword(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(dirWithHTML + "cabinet-password.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
