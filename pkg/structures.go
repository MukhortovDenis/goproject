package pkg

type User struct {
	ID         int    `json:"id"`
	First_name string `json:"firts_name"`
	Last_name  string `json:"last_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type Config struct {
	Port string `yaml:"port" env:"PORT"`
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
}
