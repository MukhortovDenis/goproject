package main

import (
	"database/sql"
	"fmt"
	"goproject/pkg"
	"log"
	"net/http"
	"os"
)

// Аутентификация
// Есть маленькая проверка на валидность, скоро заменится на js
func check(w http.ResponseWriter, r *http.Request) {
	var checkUser pkg.User
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		fmt.Fprint(w, "Не все данные введены")
	}
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	rows, err := db.Query("SELECT * FROM users WHERE login = $1", login)
	if err != nil {
		fmt.Fprint(w, "Неправильный логин")
	}
	for rows.Next() {
		err = rows.Scan(&checkUser.ID, &checkUser.First_name, &checkUser.Last_name, &checkUser.Login, &checkUser.Password)
		if err != nil {
			panic(err)
		}
	}
	if checkUser.Password == password {

		session, err := store.Get(r, "session")
		if err != nil {
			log.Print(err)
		}
		session.Values["userID"] = checkUser.ID
		session.Values["firstname"] = checkUser.First_name
		session.Values["lastname"] = checkUser.Last_name
		err = session.Save(r, w)
		if err != nil {
			log.Print(err)
		}
	} else {
		fmt.Fprint(w, "Неправильный пароль")
	}
	defer rows.Close()
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
