package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var MySecret string

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	MySecret = os.Getenv("SECRET_SAUCE")
}

func GetSecret() string {
	return MySecret
}
