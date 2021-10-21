package main

import (
	"log"

	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

// Путь к статике для рендеринга html со стороны сервера
var dirWithHTML string = "./ui/html/"

// Создание структуры, в которой подбираются данные из окружения
var configEnv = init_env()

// URI к бд
var dbConn string = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=require", configEnv.Dialect, configEnv.DataUser, configEnv.DataPass, configEnv.DataHost, configEnv.DataPort, configEnv.DataName)

//Создание хранилища куки с рандомным ключом
var store = sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))

// Все основные обработчики сервера
func mainHandle() *chi.Mux {
	// Создание go-chi роутера с доп. логированием
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
			err = tmp.Execute(w, nil)
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
			err = tmp.Execute(w, nil)
			if err != nil {
				fmt.Fprint(w, err)
			}
		})
	// Главная
	router.Get("/",
		func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "session")
			if err != nil {
				log.Print(err)
			}
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
	//Обработчики данных

	// Регистрация нового пользователя
	router.Get("/save_user", save)

	// Выход из аккаунта(удаление данных из сессии)
	router.Get("/quit", quit)

	// Аутентификация
	router.Get("/check_user", check)
	return router
}
