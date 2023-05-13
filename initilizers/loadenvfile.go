package initilizers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cann't load env files")
	}
}
