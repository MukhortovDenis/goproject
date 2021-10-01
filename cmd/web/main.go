package main

import (
	"fmt"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	cfg := Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Println(err)
	}
	path := cfg.Host + ":" + cfg.Port
	fmt.Printf("starting server at %s", path)
	err = http.ListenAndServe(path, nil)
	if err != nil {
		panic(err)
	}
	mainHandle()
}
