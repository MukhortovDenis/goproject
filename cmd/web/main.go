package main

import (
	"database/sql"
	"fmt"
	"goproject/pkg"
	"log"
	"os"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	db, err := sql.Open("postgres", pkg.DBConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if err = db.Ping(); err != nil{
		fmt.Fprintf(os.Stderr, "Unable to ping to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	handler := &pkg.Handler{
		Store:   sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32))),
		Storage: db,
	}
	// m := &autocert.Manager{
	// 	Cache:      autocert.DirCache("golang-autocert"),
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("https://stoneshop.herokuapp.com", "stoneshop.herokuapp.com", "127.0.0.1:443"),
	// }
	cfg := pkg.Config{}
	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	path := cfg.Host + ":" + cfg.Port
	server := new(pkg.Server)
	err = server.Run(path, handler.MainHandle())
	if err != nil {
		log.Fatal(err)
	}
}
