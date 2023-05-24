package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetConfigs()(DBconf, APIconf, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
	var Database DBconf
	Database.Hostname = os.Getenv("DB_HOSTNAME")
	Database.Port = os.Getenv("DB_PORT")
	Database.User = os.Getenv("DB_USER")
	Database.Password = os.Getenv("DB_PASSWORD")
	Database.Database = os.Getenv("DB_DATABASE")

	var Api APIconf
	Api.Port = os.Getenv("API_PORT")

	return Database, Api, nil
}

type APIconf struct {
	Port string
}
type DBconf struct {
	Hostname string
	Port     string
	User     string
	Password string
	Database string
}
