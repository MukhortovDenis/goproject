package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var dirWithHTML string = "./ui/html/"

func handl(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(dirWithHTML + "index.html")
	if err != nil {
		fmt.Println(err)
	}
	tmp.Execute(w, nil) // нил на энное время
}

func mainHandle() {
	http.HandleFunc("/signin",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "signin.html")
			if err != nil {
				fmt.Println(err)
			}
			tmp.Execute(w, nil) // нил на энное время

		})

	http.HandleFunc("/signup",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "signup.html")
			if err != nil {
				fmt.Println(err)
			}
			tmp.Execute(w, nil) // нил на энное время
			fmt.Fprintln(w, "Здесб рега:", r.URL.String())
		})
	http.HandleFunc("/save_user",
		func(w http.ResponseWriter, r *http.Request) {
			login := r.FormValue("login")
			password := r.FormValue("password")
			db, err := sql.Open("mysql", "mysql:123@tcp(127.0.0.1:3306)/stoneshop")
			if err != nil {
				panic(err)
			}
			insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`login`, `password`) VALUES('%s', '%s')", login, password))
			if err != nil {
				panic(err)
			}
			defer insert.Close()
			defer db.Close()
			http.Redirect(w, r, "/", http.StatusSeeOther)
		})
	http.HandleFunc("/", handl)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
