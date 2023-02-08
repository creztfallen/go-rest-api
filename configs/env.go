package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// load envs
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_URI")
}
