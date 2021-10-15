package main

import (
	"goproject/pkg"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init_env() *pkg.ConfigEnv {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	key, _ := os.LookupEnv("KEY_STORE")
	dialect, _ := os.LookupEnv("DIALECT")
	data_user, _ := os.LookupEnv("DATABASE_USER")
	data_pass, _ := os.LookupEnv("DATABASE_PASSWORD")
	data_host, _ := os.LookupEnv("DATABASE_HOST")
	data_port, _ := os.LookupEnv("DATABASE_PORT")
	data_name, _ := os.LookupEnv("DATABASE_NAME")

	configEnv := &pkg.ConfigEnv{
		KeyStore: key,
		Dialect:  dialect,
		DataUser: data_user,
		DataPass: data_pass,
		DataHost: data_host,
		DataPort: data_port,
		DataName: data_name,
	}
	return configEnv
}
