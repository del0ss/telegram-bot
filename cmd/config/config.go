package config

import (
	"os"
	"strconv"
)

type TelegramConfig struct {
	TelegramBotAPI   string
	DebugTelegramBot bool
	AdminID          int64
}

type RedisConfig struct {
	RedisHost     string
	RedisPassword string
	RedisDB       int
}

type ExcelConfig struct {
	ExcelFileName string
}

type Config struct {
	Telegram TelegramConfig
	Redis    RedisConfig
	Excel    ExcelConfig
}

func New() *Config {
	return &Config{
		Telegram: TelegramConfig{
			TelegramBotAPI:   getEnv("telegram_bot_api", ""),
			DebugTelegramBot: getEnvAsBool("debug_telegram_bot", false),
			AdminID:          getEnvAsInt64("admin", 0),
		},
		Redis: RedisConfig{
			RedisHost:     getEnv("redis_host", ""),
			RedisPassword: getEnv("redis_password", ""),
			RedisDB:       getEnvAsInt("redis_db", 0),
		},
		Excel: ExcelConfig{
			ExcelFileName: getEnv("excel_file_name", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsInt64(name string, defaultVal int64) int64 {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return int64(value)
	}
	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultVal
}
