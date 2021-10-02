package main

import (
	"context"
	"fmt"
	"goproject/pkg"
	"html/template"
	"net/http"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v4"
)

var user pkg.User
var connStr string = "postgres://nigger:nigger@localhost:5432"

// Путь до шаблоном, мб быстрее на пару мгновений, если буду указывать не через переменную
var dirWithHTML string = "./ui/html/"

// Подключение к локальной бд, где после регистрации новый пользователь добавляет новую запись
func save(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		fmt.Fprint(w, "Не все данные введены")
	}
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err := conn.QueryRow(context.Background(), "INSERT INTO users (login, password) VALUES ($1, $2)", login, password)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Проверка, есть ли запись пользователя в бд по логину и паролю(пока локально)
func check(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		fmt.Fprint(w, "Не все данные введены")
	}
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// search, err := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `login`='%s'", login))
	// if err != nil {
	// 	fmt.Fprint(w, "Неправильный логин")
	// }
	// for search.Next() {
	// 	err = search.Scan(&user.ID, &user.Login, &user.Password)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if password != user.Password {
	// 		fmt.Fprint(w, "Неправильный пароль")
	// 	}
	// }

	// defer search.Close()
	defer db.Close(context.Background())
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
