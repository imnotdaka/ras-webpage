package config

type Config struct {
	DB DB
}

type DB struct {
	User     string
	Password string
	Database string
	Ip       string
	Port     string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	cfg.DB.User = "root"
	cfg.DB.Password = "123"
	cfg.DB.Database = "todoapp"
	cfg.DB.Ip = "localhost"
	cfg.DB.Port = "3306"

	return cfg, nil
}
