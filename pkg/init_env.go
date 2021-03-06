package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Работа с переменным окружением, а именно с файлом .env
func init_env() ConfigEnv {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	dialect, _ := os.LookupEnv("DIALECT")
	data_user, _ := os.LookupEnv("DATABASE_USER")
	data_pass, _ := os.LookupEnv("DATABASE_PASSWORD")
	data_host, _ := os.LookupEnv("DATABASE_HOST")
	data_port, _ := os.LookupEnv("DATABASE_PORT")
	data_name, _ := os.LookupEnv("DATABASE_NAME")

	configEnv := ConfigEnv{
		Dialect:  dialect,
		DataUser: data_user,
		DataPass: data_pass,
		DataHost: data_host,
		DataPort: data_port,
		DataName: data_name,
	}
	return configEnv
}
