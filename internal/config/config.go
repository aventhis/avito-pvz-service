package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config содержит все настройки приложения
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	GRPC     GRPCConfig
	Metrics  MetricsConfig
}

// DatabaseConfig настройки базы данных
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	DSN      string // Строка подключения к базе данных
}

// ServerConfig настройки HTTP-сервера
type ServerConfig struct {
	Port int
	Env  string
}

// GRPCConfig настройки gRPC-сервера
type GRPCConfig struct {
	Port int
}

// MetricsConfig настройки Prometheus
type MetricsConfig struct {
	Port int
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	appPort, err := strconv.Atoi(getEnv("APP_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid APP_PORT: %w", err)
	}

	grpcPort, err := strconv.Atoi(getEnv("GRPC_PORT", "3000"))
	if err != nil {
		return nil, fmt.Errorf("invalid GRPC_PORT: %w", err)
	}

	metricsPort, err := strconv.Atoi(getEnv("METRICS_PORT", "9000"))
	if err != nil {
		return nil, fmt.Errorf("invalid METRICS_PORT: %w", err)
	}

	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     dbPort,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Name:     getEnv("DB_NAME", "avito_pvz"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	// Формируем строку подключения к базе данных
	dbConfig.DSN = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password,
		dbConfig.Name, dbConfig.SSLMode,
	)

	config := &Config{
		Database: dbConfig,
		Server: ServerConfig{
			Port: appPort,
			Env:  getEnv("APP_ENV", "development"),
		},
		GRPC: GRPCConfig{
			Port: grpcPort,
		},
		Metrics: MetricsConfig{
			Port: metricsPort,
		},
	}

	return config, nil
}

// Получить переменную окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 