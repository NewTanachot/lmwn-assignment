package configuration

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	Port        = "PORT"
	CovidApiUrl = "LMWN_COVID_ENDPOINT"
)

func init() {
	godotenv.Load()
}

func GetCovidUrl() string {
	defaultUrl := "https://static.wongnai.com/devinterview/covid-cases.json"
	url := os.Getenv(CovidApiUrl)

	if url == "" {
		return defaultUrl
	}

	return url
}

func GetEnvPort() string {
	defaultPort := "8080"
	port := os.Getenv(Port)

	if port == "" {
		port = defaultPort
	}

	return fmt.Sprintf(":%s", port)
}
