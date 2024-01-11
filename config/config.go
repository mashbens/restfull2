package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Driver  string
	DB_Host string
	DB_Port string
	DB_User string
	DB_Pass string
	DB_Name string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &AppConfig{
		Driver:  os.Getenv("DRIVER"),
		DB_Host: os.Getenv("DB_HOST"),
		DB_Port: os.Getenv("DB_PORT"),
		DB_User: os.Getenv("DB_USER"),
		DB_Pass: os.Getenv("DB_PASS"),
		DB_Name: os.Getenv("DB_NAME"),
	}
}

type ServerConfig struct {
	Host string
	Port string
}
