package configs

import "github.com/caarlos0/env/v6"

var (
	conf *Config
)

type Config struct {
	GinPort       string `env:"GIN_PORT" envDefault:"8080"`
	MySQLHost     string `env:"MYSQL_HOST"`
	MySQLPort     int    `env:"MYSQL_PORT" envDefault:"3306"`
	MySQLUser     string `env:"MYSQL_USER"`
	MySQLPassword string `env:"MYSQL_PASSWORD"`
	MySQLDatabase string `env:"MYSQL_DATABASE"`
}

func Init() {
	conf = &Config{}
	if err := env.Parse(conf); err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return conf
}
