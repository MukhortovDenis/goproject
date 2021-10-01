package main

import (
	"database/sql"
	"fmt"
	"goproject/pkg"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var user pkg.User

// Путь до шаблоном, мб быстрее на пару мгновений, если буду указывать не через переменную
var dirWithHTML string = "./ui/html/"

// Подключение к локальной бд, где после регистрации новый пользователь добавляет новую запись
func save(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		fmt.Fprint(w, "Не все данные введены")
	}
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

//Проверка, есть ли запись пользователя в бд по логину и паролю(пока локально)
func check(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		fmt.Fprint(w, "Не все данные введены")
	}
	db, err := sql.Open("mysql", "mysql:123@tcp(127.0.0.1:3306)/stoneshop")
	if err != nil {
		panic(err)
	}
	search, err := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `login`='%s'", login))
	if err != nil {
		fmt.Fprint(w, "Неправильный логин")
	}
	for search.Next() {
		err = search.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			panic(err)
		}
		if password != user.Password {
			fmt.Fprint(w, "Неправильный пароль")
		}
	}

	defer search.Close()
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Страницы, которые отображаются у пользователей
// Пока нет готового дизайна, новые делать не буду((
func mainHandle() {
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
}
