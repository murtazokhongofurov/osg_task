package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HttpPort         string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresHost     string
	PostgresPort     string
	LogLevel         string
	PgxPoolSize      int
	SigninKey        string
	AuthConfigPath   string
	CsvFilePath      string
}

func Load() Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Error connect .env loading", err.Error())
	}
	c := Config{}

	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "http_port"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "host"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "port"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "username"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "password"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "dbname"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.PgxPoolSize = cast.ToInt(getOrReturnDefault("PGXPOOL_MAXSIZE", 2))
	c.SigninKey = cast.ToString(getOrReturnDefault("SIGNIN_KEY", "signinkey"))
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_FILE_PATH", "./internal/pkg/config/auth.conf"))
	c.CsvFilePath = cast.ToString(getOrReturnDefault("CSV_FILE_PATH", "./internal/pkg/config/roles.csv"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
