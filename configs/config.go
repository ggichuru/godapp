package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// constant highlighting the env source file.
const envloc = ".env"

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loadind env from: %s \n %v", envloc, err)
	}
}
