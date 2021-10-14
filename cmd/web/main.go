package main

import (
	"context"
	"goproject/pkg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/sessions"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	router := mainHandle()
	// m := &autocert.Manager{
	// 	Cache:      autocert.DirCache("golang-autocert"),
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("https://stoneshop.herokuapp.com", "stoneshop.herokuapp.com", "127.0.0.1:443"),
	// }
	cfg := pkg.Config{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	path := cfg.Host + ":" + cfg.Port
	server := &http.Server{
		Addr:    path,
		Handler: router,
		// TLSConfig: m.TLSConfig(),
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	session := &sessions.Session{}
	if session.Values != nil {
		session.Options.MaxAge = 0
		session.Values = nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
