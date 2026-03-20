package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load(".env")
	godotenv.Load("../history-service/.env")
}

type GRPCConfig struct {
	Addr string
}

func NewGRPCConfig() *GRPCConfig {
	return &GRPCConfig{
		Addr: getString("GRPC_ADDR", ":9090"),
	}
}

type DatabaseConfig struct {
	Url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: getRequiredString("POSTGRES_URL"),
	}
}

type LoggerConfig struct {
	File   string
	Level  int
	Format string
}

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		File:   getString("LOG_FILE", ""),
		Level:  getInt("LOG_LEVEL", 0),
		Format: getString("LOG_FORMAT", "json"),
	}

}

func getString(key, defaultValue string) string {
	str := os.Getenv(key)
	if str == "" {
		return defaultValue
	}
	return str
}

func getRequiredString(key string) string {
	str := os.Getenv(key)
	if str == "" {
		log.Fatalf("обязательная переменная окружения не задана: %s", key)
	}
	return str
}

func getInt(key string, defaultValue int) int {
	str := os.Getenv(key)
	if str == "" {
		return defaultValue
	}
	var i int
	if _, err := fmt.Sscanf(str, "%d", &i); err != nil {
		return defaultValue
	}
	return i
}
