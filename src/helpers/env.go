package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string, defValue string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("env not found %s", err)
	}
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defValue
}
