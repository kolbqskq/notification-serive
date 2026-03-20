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
	godotenv.Load("../.env.shared", "notification-worker/.env")

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
	Topic   string
	GroupID string
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Brokers: getStringSlice("KAFKA_BROKERS"),
		Topic:   getRequiredString("KAFKA_TOPIC"),
		GroupID: getRequiredString("KAFKA_GROUP_ID"),
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

type DatabaseConfig struct {
	Url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: getRequiredString("DATABASE_URL"),
	}
}

type TelegramConfig struct {
	Token  string
	ChatID int64
}

func NewTelegramConfig() *TelegramConfig {
	return &TelegramConfig{
		Token:  getRequiredString("TELEGRAM_BOT_TOKEN"),
		ChatID: getRequiredInt64("TELEGRAM_CHAT_ID"),
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

func getRequiredInt64(key string) int64 {
	str := os.Getenv(key)
	int64Str, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalf("обязательная переменная окружения не задана: %s", key)
	}
	return int64Str
}

func getStringSlice(key string) []string {
	val := getRequiredString(key)
	return strings.Split(val, ",")
}
