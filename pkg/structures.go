package pkg

type User struct {
	ID         int    `json:"-"`
	First_name string `json:"firstname"`
	// Last_name  string `json:"last_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserDB struct {
	ID         int
	First_name string
	// Last_name  string `json:"last_name"`
	Login    string
	Password string
}
type Stone struct {
	ID          int
	Name        string
	URL         string
	Description string
	Price       int
	Rare        string
}

type Config struct {
	Port string `yaml:"port" env:"PORT"`
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`
}
type ConfigEnv struct {
	KeyStore string
	Dialect  string
	DataUser string
	DataPass string
	DataHost string
	DataPort string
	DataName string
}
