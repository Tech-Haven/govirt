package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Env loaded...")
}

func EnvMongoUri() string {
	return os.Getenv("MONGO_URI")
}

func EnvOvirtUrl() string {
	return os.Getenv("OVIRT_URL")
}