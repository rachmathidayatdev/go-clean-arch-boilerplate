package config

import "os"

//Config struct
type Config struct {
	DB *DBConfig
}

//DBConfig struct
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Database string
	Host     string
	Port     string
	Charset  string
}

//GetConfigDB function
func GetConfigDB() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASS"),
			Database: os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Charset:  "utf8",
		},
	}
}
