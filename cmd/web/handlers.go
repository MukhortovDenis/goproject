package main

import (
	"database/sql"
	"fmt"
	"goproject/pkg"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Путь до шаблоном, мб быстрее на пару мгновений, если буду указывать не через переменную
var dirWithHTML string = "./ui/html/"

// Подключение к локальной бд, где после регистрации новый пользователь добавляет новую запись
func save(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		fmt.Println("Не все данные введены")
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
		fmt.Println("Не все данные введены")
	}
	db, err := sql.Open("mysql", "mysql:123@tcp(127.0.0.1:3306)/stoneshop")
	if err != nil {
		panic(err)
	}
	search, err := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `login`='%s'", login))
	if err != nil {
		fmt.Println("Неправильный логин")
	}
	var user pkg.User
	for search.Next() {
		err = search.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			panic(err)
		}
		if password != user.Password {
			fmt.Println("Неправильный пароль")
		}
	}

	defer search.Close()
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Страницы, которые отображаются у пользователей
// Пока нет готового дизайна, новые делать не буду((
func mainHandle() {
	// Регистрация
	http.HandleFunc("/signin",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "signin.html")
			if err != nil {
				fmt.Println(err)
			}
			tmp.Execute(w, nil) // нил на энное время

		})
	//Вход
	http.HandleFunc("/signup",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "signup.html")
			if err != nil {
				fmt.Println(err)
			}
			tmp.Execute(w, nil) // нил на энное время
		})
	// Главная
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			tmp, err := template.ParseFiles(dirWithHTML + "index.html")
			if err != nil {
				fmt.Println(err)
			}
			tmp.Execute(w, nil) // нил на энное время)
		})
	// То, что пользователь не увидит, пока только сохранение и проверка записи в бд
	http.HandleFunc("/save_user", save)

	http.HandleFunc("/check_user", check)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
