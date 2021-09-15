package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handl(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("index.html") // разобраться потом, хтмл пока поваляется в этом каталоге
	if err != nil {
		fmt.Println("хтмл не подхатился")
	}
	tmp.Execute(w, nil) // нил на энное время
}

func mainHandle() {
	http.HandleFunc("/signin",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Здесь лога:", r.URL.String())
		})

	http.HandleFunc("/signup",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Здесб рега:", r.URL.String())
		})

	http.HandleFunc("/", handl)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
