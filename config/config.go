package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	DB  *DBConfig
}

func LoadConfig() Config {
	envPath := filepath.Join("/", ".env.local")
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("WARN: Error loading .env file: %v", err)
	}

	conf := Config{}

	// DB setting
	conf.DB = &DBConfig{}
	conf.DB.Type = os.Getenv("DB_TYPE")
	conf.DB.Host = os.Getenv("DB_HOST")
	conf.DB.Name = os.Getenv("DB_NAME")
	conf.DB.UserName = os.Getenv("DB_USERNAME")
	conf.DB.Password = os.Getenv("DB_PASSWORD")

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	conf.DB.Port = dbPort

	dbPoolConns, err := strconv.Atoi(os.Getenv("DB_POOL_CONNS"))
	if err != nil {
		panic(err)
	}

	conf.DB.PoolConns = dbPoolConns

	return conf
}
