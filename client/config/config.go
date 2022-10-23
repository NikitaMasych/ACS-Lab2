package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerHost string
	ServerPort string
	ServerType string
)

func init() {
	loadEnv()
	setupVariables()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func setupVariables() {
	ServerHost = os.Getenv("SERVER_HOST")
	ServerPort = os.Getenv("SERVER_PORT")
	ServerType = os.Getenv("SERVER_TYPE")
}
