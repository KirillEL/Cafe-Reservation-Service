package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	JwtSecret  string
}

var Env *Environment

func getEnv(key string, required bool) string {
	value, ok := os.LookupEnv(key)
	if !ok && required {
		log.Fatalf("")
	}
	return value
}

func LoadEnvironment() {
	if Env == nil {
		Env = new(Environment)
	}
	Env.DBName = getEnv("DB_NAME", true)
	Env.DBHost = getEnv("DB_HOST", true)
	Env.DBPassword = getEnv("DB_PASSWORD", true)
	Env.DBPort = getEnv("DB_PORT", true)
	Env.DBUser = getEnv("DB_USER", true)
	Env.JwtSecret = getEnv("JWT_SECRET", true)

}

func LoadEnvFile(file string) {
	if err := godotenv.Load(file); err != nil {
		fmt.Printf("Error loading env file: %s", err)
	}
	LoadEnvironment()
}
