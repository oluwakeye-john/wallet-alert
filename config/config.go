package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	error := godotenv.Load()
	if error != nil {
		log.Fatalf(error.Error())
	}
}

func MustGetEnv(s string) string {
	k := os.Getenv(s)
	if k == "" {
		log.Fatalf("variable %s is not set", s)
	}
	return k
}

func GetEnv(s string) string {
	return os.Getenv(s)
}
