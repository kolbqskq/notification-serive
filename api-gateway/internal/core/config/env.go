package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load(".env")
	godotenv.Load("../.env.shared", "api-gateway/.env")

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

type RedisConfig struct {
	RedisAddr string
}

func NewAddrConfig() *RedisConfig {
	return &RedisConfig{
		RedisAddr: getRequiredString("REDIS_ADDR"),
	}
}

type KafkaConfig struct {
	Brokers []string
	Topic string
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Brokers: getStringSlice("KAFKA_BROKERS"),
		Topic: getRequiredString("KAFKA_TOPIC"),
	}
}

type ServiceConfig struct {
	Name string
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		Name: getRequiredString("SERVICE_NAME"),
	}
}

type AppConfig struct {
	Addr string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Addr: getString("APP_ADDR", ":8081"),
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
	intStr, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return intStr
}

func getStringSlice(key string) []string {
	val := getRequiredString(key)
	return strings.Split(val, ",")
}
