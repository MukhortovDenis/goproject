package pkg

import (
	"bytes"
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
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Print(err)
	}
	if newUser.Login == "" || newUser.Password == "" || newUser.First_name == "" {
		Error := new(Error)
		Error.NewErrorMessage("Пустые поля")
		body := new(bytes.Buffer)
		err = json.NewEncoder(body).Encode(Error)
		if err != nil {
			log.Print(err)
		}
		fmt.Fprint(w, body)
	}
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	var checkLogin string
	row, err := db.Query("SELECT login FROM users WHERE login = $1", newUser.Login)
	if err != nil {
		fmt.Fprint(w, err)
	}
	for row.Next() {
		err = row.Scan(&checkLogin)
		if err != nil {
			log.Fatal(err)
		}
	}
	if newUser.Login == checkLogin {
		Error := new(Error)
		Error.NewErrorEmail(true)
		body := new(bytes.Buffer)
		err = json.NewEncoder(body).Encode(Error)
		if err != nil {
			log.Print(err)
		}
		fmt.Fprint(w, body)

	} else {
		var userid int
		err = db.QueryRow(`INSERT INTO users (firstname, login, password) VALUES ($1, $2, $3) RETURNING id`, newUser.First_name, newUser.Login, newUser.Password).Scan(&userid)
		if err != nil {
			fmt.Fprint(w, err)
		}
		fmt.Fprint(w, "{}")
	}
	defer row.Close()
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
			log.Fatal(err)
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
		if err != nil {
			log.Print(err)
		}
		fmt.Fprint(w, "{}")
	} else {
		Error := new(Error)
		Error.NewErrorPass(true)
		body := new(bytes.Buffer)
		err = json.NewEncoder(body).Encode(Error)
		if err != nil {
			log.Print(err)
		}
		fmt.Fprint(w, body)
	}
	defer rows.Close()
	defer db.Close()
	defer r.Body.Close()
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

func (h *Handler) resetCabinetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := store.Get(r, "session")
	if err != nil {
		log.Print(err)
	}
	var data newData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var checkLogin string
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	row, err := db.Query("SELECT login FROM users WHERE login = $1", data.NewEmail)
	if err != nil {
		fmt.Fprint(w, err)
	}
	for row.Next() {
		err = row.Scan(&checkLogin)
		if err != nil {
			log.Fatal(err)
		}
	}
	if data.NewEmail == checkLogin {
		Error := new(Error)
		Error.NewErrorEmail(true)
		body := new(bytes.Buffer)
		err = json.NewEncoder(body).Encode(Error)
		if err != nil {
			log.Print(err)
		}
		fmt.Fprint(w, body)

	} else {
		oldEmail := session.Values["email"]
		log.Println(oldEmail)
		row, err = db.Query("UPDATE users SET firstname = $1 , login = $2 WHERE login = $3", data.NewFirstName, data.NewEmail, oldEmail)
		if err != nil {
			log.Println(err)
		}
		log.Println(session.Values["firstname"], session.Values["email"])
		session.Values["firstname"] = data.NewFirstName
		session.Values["email"] = data.NewEmail
		log.Println(session.Values["firstname"], session.Values["email"])
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprint(w, "{}")
		defer db.Close()
		defer row.Close()
		defer r.Body.Close()
	}
}
