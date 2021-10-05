package main

// Отрефакторить
import (
	"database/sql"

	"fmt"
	"goproject/pkg"
	"html/template"
	"net/http"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

var user pkg.User
var connStr string = "postgres://postgres:123@localhost:5432/stone_shop?sslmode=disable"

// Путь до шаблоном, мб быстрее на пару мгновений, если буду указывать не через переменную
var dirWithHTML string = "./ui/html/"

// Подключение к локальной бд, где после регистрации новый пользователь добавляет новую запись
func save(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		fmt.Fprint(w, "Не все данные введены")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	var newUser pkg.User
	err = db.QueryRow(`INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id`, login, password).Scan(&newUser.ID)
	if err != nil {
		fmt.Fprint(w, err)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Проверка, есть ли запись пользователя в бд по логину и паролю(пока локально)
func check(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		fmt.Fprint(w, "Не все данные введены")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	var newUser pkg.User
	err = db.QueryRow("SELECT * FROM users WHERE login = $1", login).Scan(&newUser.ID, &newUser.Login, &newUser.Password)
	if err != nil {
		fmt.Fprint(w, "Неправильные данные")
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Страницы, которые отображаются у пользователей
// Пока нет готового дизайна, новые делать не буду((
func mainHandle() {
	cfg := Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Println(err)
	}
	// Отслеживание сервером статических файлов
	fsForCss := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fsForCss))
	//
	fsForImg := http.FileServer(http.Dir("./ui/img/"))
	http.Handle("/img/", http.StripPrefix("/img", fsForImg))
	// Регистрация
	http.HandleFunc("/signin",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "signin.html")
			if err != nil {
				fmt.Println(err)
			}
			err = tmp.Execute(w, user) // нил на энное время
			if err != nil {
				fmt.Fprint(w, err)
			}

		})
	//Вход
	http.HandleFunc("/signup",
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
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "index.html")
			if err != nil {
				fmt.Println(err)
			}
			err = tmp.Execute(w, user) // нил на энное время)
			if err != nil {
				fmt.Fprint(w, err)
			}
		})
	// То, что пользователь не увидит, пока только сохранение и проверка записи в бд
	http.HandleFunc("/save_user", save)

	http.HandleFunc("/check_user", check)
	path := cfg.Host + ":" + cfg.Port
	err = http.ListenAndServe(path, nil)
	if err != nil {
		panic(err)
	}
}
