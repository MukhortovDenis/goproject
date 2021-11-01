package main

import (
	"context"
	"goproject/pkg"
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	handler := new(pkg.Handler)
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
	server := new(pkg.Server)
	err = server.Run(path, handler.MainHandle())
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
