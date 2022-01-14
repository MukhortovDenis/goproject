package main

import (
	"goproject/pkg"
	"log"

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
}
