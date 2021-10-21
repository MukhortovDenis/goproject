package main

import (
	"database/sql"
	"fmt"
	"goproject/pkg"
	"net/http"
	"os"
)

// Сохранение данных о пользователе в бд
func save(w http.ResponseWriter, r *http.Request) {
	var newUser pkg.User
	newUser.First_name = r.FormValue("firstname")
	newUser.Last_name = r.FormValue("lastname")
	newUser.Login = r.FormValue("login")
	newUser.Password = r.FormValue("password")
	passwordCheck := r.FormValue("password-check")
	if newUser.Login == "" || newUser.Password == "" || newUser.First_name == "" || newUser.Last_name == "" || passwordCheck == "" {
		fmt.Fprint(w, "Не все данные введены")
	}
	if newUser.Password != passwordCheck {
		fmt.Fprint(w, "Пароли не сходятся")
	}
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	var userid int
	err = db.QueryRow(`INSERT INTO users (firstname, lastname, login, password) VALUES ($1, $2, $3, $4) RETURNING id`, newUser.First_name, newUser.Last_name, newUser.Login, newUser.Password).Scan(&userid)
	if err != nil {
		fmt.Fprint(w, err)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
