package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAdress   string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: getENV("LOCAL_HOST", "http://localhost"),
		Port:       getENV("PORT", ":8000"),
		DBUser:     getENV("DB_USER", "root"),
		DBPassword: getENV("DB_PASSWORD", "C0ntr4s3n4-"),
		DBAdress:   fmt.Sprintf("%s%s", getENV("DB_HOST", "127.0.0.1"), getENV("DB_PORT", ":3306")),
		DBName:     getENV("DB_NAME", "go_blog"),
	}
}

func getENV(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
