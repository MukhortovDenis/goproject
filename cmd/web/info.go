package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func info(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
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
