package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Port string `env:"APP_PORT"`
		Env  string `env:"APP_ENV"`
		Name string `env:"APP_NAME"`
	}
	Postgres struct {
		Host string `env:"POSTGRES_HOST"`
		Port string `env:"POSTGRES_PORT"`
		DB   string `env:"POSTGRES_DB"`
		User string `env:"POSTGRES_USER"`
		Pass string `env:"POSTGRES_PASS"`
	}
	Redis struct {
		Host string `env:"REDIS_HOST"`
		Port string `env:"REDIS_PORT"`
	}
}

func Loadenv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{}

	cfg.App.Port = os.Getenv("APP_PORT")
	cfg.App.Env = os.Getenv("APP_ENV")
	cfg.App.Name = os.Getenv("APP_NAME")

	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.DB = os.Getenv("POSTGRES_DB")
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Pass = os.Getenv("POSTGRES_PASS")

	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Port = os.Getenv("REDIS_PORT")

	return cfg
}
