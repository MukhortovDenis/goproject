package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func (h *Handler) save(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Print(err)
	}
	if newUser.Login == "" || newUser.Password == "" || newUser.First_name == "" {
		fmt.Fprint(w, "Не все данные введены")
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
	defer r.Body.Close()
}

func (h *Handler) check(w http.ResponseWriter, r *http.Request) {
	var CheckUser User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&CheckUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	rows, err := db.Query("SELECT * FROM users WHERE login = $1", CheckUser.Login)
	if err != nil {
		fmt.Fprint(w, "Неправильный логин")
	}
	var user UserDB
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.First_name, &user.Login, &user.Password)
		if err != nil {
			panic(err)
		}
	}
	if CheckUser.Password == user.Password {
		session, err := store.Get(r, "session")
		if err != nil {
			log.Print(err)
		}
		session.Values["userID"] = user.ID
		session.Values["firstname"] = user.First_name
		session.Values["email"] = user.Login
		err = session.Save(r, w)
		if err != nil {
			log.Print(err)
		}
	} else {
		fmt.Fprint(w, "Неправильный пароль")
	}
	defer rows.Close()
	defer db.Close()
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

// func (h *Handler) checkPost(w http.ResponseWriter, r *http.Request) {
// 	session, err := store.Get(r, "session")
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	for {
// 		login := session.Values["userID"]
// 		if login != nil {
// 			break
// 		}
// 		http.Redirect(w, r, "/", http.StatusSeeOther)

// 	}
// }
