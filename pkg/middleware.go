package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func (h *Handler) save(w http.ResponseWriter, r *http.Request) {
	var newUser User
	newUser.First_name = r.FormValue("firstname")
	newUser.Login = r.FormValue("login")
	newUser.Password = r.FormValue("password")
	passwordCheck := r.FormValue("password-check")
	if newUser.Login == "" || newUser.Password == "" || newUser.First_name == "" || passwordCheck == "" {
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
	err = db.QueryRow(`INSERT INTO users (firstname, login, password) VALUES ($1, $2, $3) RETURNING id`, newUser.First_name, newUser.Login, newUser.Password).Scan(&userid)
	if err != nil {
		fmt.Fprint(w, err)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) check(w http.ResponseWriter, r *http.Request) {
	var checkUser User
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
		err = rows.Scan(&checkUser.ID, &checkUser.First_name, &checkUser.Login, &checkUser.Password)
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
		session.Values["email"] = checkUser.Login
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

func (h *Handler) quit(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Print(err)
	}
	for i := range session.Values {
		session.Values[i] = nil
	}
	err = session.Save(r, w)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
