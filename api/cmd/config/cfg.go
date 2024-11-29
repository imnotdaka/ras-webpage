package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB            DB
	MPAccessToken string
	JWTSecret     string
}

type DB struct {
	User     string
	Password string
	Database string
	Ip       string
	Port     string
}

func NewConfig() (*Config, error) {
	dir := "./cmd/config/.env"
	err := godotenv.Load(dir)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	cfg.DB.User = os.Getenv("USERENV")
	cfg.DB.Password = os.Getenv("PASSWORDENV")
	cfg.DB.Database = os.Getenv("DATABASEENV")
	cfg.DB.Ip = os.Getenv("IPENV")
	cfg.DB.Port = os.Getenv("PORTENV")

	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.MPAccessToken = os.Getenv("ACCESS_TOKEN")

	return cfg, nil
}
