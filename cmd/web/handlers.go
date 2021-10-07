package main

// Отрефакторить, БЛЯТЬ ГИТХАБ ЛАГАЕТ
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

// Путь до шаблоном, мб быстрее на пару мгновений, если буду указывать не через переменную
var dirWithHTML string = "./ui/html/"
var connStr string = "postgres://kfireyqrkgozaa:31b2140dfdba297c412bda66a9db337c91a8729b17a9791bea82c934ff095d4c@ec2-34-249-247-7.eu-west-1.compute.amazonaws.com:5432/d900njt9tj61n8?sslmode=require"
var userDefault pkg.User

// Подключение к локальной бд, где после регистрации новый пользователь добавляет новую запись
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
	db, err := sql.Open("postgres", connStr)
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
	db, err := sql.Open("postgres", connStr)
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
		if password != checkUser.Password {
			fmt.Fprint(w, "Неправильный пароль")
			checkUser = pkg.User{}
		}
		userDefault = checkUser
		defer rows.Close()
		defer db.Close()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
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
			err = tmp.Execute(w, nil) // нил на энное время
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
			err = tmp.Execute(w, userDefault) // нил на энное время)
			if err != nil {
				fmt.Fprint(w, userDefault)
			}
		})
	// То, что пользователь не увидит, пока только сохранение и проверка записи в бд
	http.HandleFunc("/save", save)

	http.HandleFunc("/check_user", check)
	path := cfg.Host + ":" + cfg.Port
	err = http.ListenAndServe(path, nil)
	if err != nil {
		panic(err)
	}
}
