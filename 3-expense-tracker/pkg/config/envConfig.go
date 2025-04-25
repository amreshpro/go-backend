package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	log.Print(".env file loaded successfully")
}


func GetEnv(key string, fallback string) string {
	value, found := os.LookupEnv(key)

	if !found || value == "" {
		log.Printf("[WARN] %s not found or empty, using fallback: %s", key, fallback)
		return fallback
	}

	return value
}



func GetEnvPort(key string, fallback string) string {
	port := GetEnv(key, fallback)
	if port != "" && port[0] != ':' {
		port = ":" + port
	}
	return port
}



func GetEnvInt(key string, fallback int) int {
	valueStr, found := os.LookupEnv(key)

	if !found || valueStr == "" {
		log.Printf("[WARN] %s not found in env, using fallback: %d", key, fallback)
		return fallback
	}

	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("[WARN] Failed to convert %s='%s' to int, using fallback: %d", key, valueStr, fallback)
		return fallback
	}

	return valueInt
}