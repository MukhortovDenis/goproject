package main

// Отрефакторить
import (
	"database/sql"
	"goproject/pkg"
	"log"

	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

var dirWithHTML string = "./ui/html/"
var configEnv = init_env()
var dbConn string = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=require", configEnv.Dialect, configEnv.DataUser, configEnv.DataPass, configEnv.DataHost, configEnv.DataPort, configEnv.DataName)
var store = sessions.NewCookieStore([]byte(configEnv.KeyStore))

//Функция выхода
func quit(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
		session.Values["userID"] = checkUser.ID
		session.Values["firstname"] = checkUser.First_name
		session.Values["lastname"] = checkUser.Last_name
		err = session.Save(r, w)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Fprint(w, "Неправильный пароль")
	}
	defer rows.Close()
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func mainHandle() *chi.Mux {
	router := NewRouter()
	// Отслеживание сервером статических файлов
	fileServer(router)
	// Регистрация
	router.Get("/signin",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "signin.html")
			if err != nil {
				fmt.Println(err)
			}
			err = tmp.Execute(w, nil) // нил на энное время
			if err != nil {
				fmt.Fprint(w, err)
			}

		})
	//Вход
	router.Get("/signup",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "signup.html")
			if err != nil {
				fmt.Println(err)
			}
			err = tmp.Execute(w, nil) // нил на энное время
			if err != nil {
				fmt.Fprint(w, err)
			}
		})
	// Главная
	router.Get("/",
		func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "session")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(dbConn)

			firstname := session.Values["firstname"]
			lastname := session.Values["lastname"]
			block := map[string]interface{}{
				"firstname":  firstname,
				"lastname":   lastname,
				"show_block": true,
			}
			if firstname == nil || lastname == nil {
				block["show_block"] = false
			}
			tmp, err := template.ParseFiles(dirWithHTML + "index.html")
			if err != nil {
				fmt.Println(err)
			}
			err = tmp.ExecuteTemplate(w, "index", block)
			if err != nil {
				log.Fatal(err)
			}
		})
	// То, что пользователь не увидит, пока только сохранение и проверка записи в бд
	router.Get("/save_user", save)
	router.Get("/quit", quit)
	router.Get("/check_user", check)
	return router
}
