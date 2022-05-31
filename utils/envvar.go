package utils

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

//EnvVar function is for read .env file
func EnvInit() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	myEnv, _ := godotenv.Read(".env")
	for key := range myEnv {
		fmt.Println("Reading dotenv config: " + key + "=" + myEnv[key])
	}
}
