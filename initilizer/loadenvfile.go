package initilizer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error while loading env file")
	}
}
