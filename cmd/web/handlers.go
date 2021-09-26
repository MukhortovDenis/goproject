package main

import (
	"database/sql"
	"fmt"
	"goproject/pkg"
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
func save(w http.ResponseWriter, r *http.Request) {
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
}
func check(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	db, err := sql.Open("mysql", "mysql:123@tcp(127.0.0.1:3306)/stoneshop")
	if err != nil {
		panic(err)
	}
	search, err := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `login`='%s'", login))
	if err != nil {
		panic(err)
	}
	var user pkg.User
	for search.Next() {
		err = search.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			panic(err)
		}
		if password != user.Password {
			panic("Пароль не тот")
		}
	}

	defer search.Close()
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
		})
	http.HandleFunc("/save_user", save)

	http.HandleFunc("/check_user", check)

	http.HandleFunc("/", handl)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
