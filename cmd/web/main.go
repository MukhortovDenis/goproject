package main

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	cfg := Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Println(err)
	}
	mainHandle()
}
