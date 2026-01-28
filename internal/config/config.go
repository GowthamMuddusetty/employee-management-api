package config

import (
	"errors"
	"os"
	"time"
)

type Config struct {
	DBHost     string        `json:"db_host"`
	DBPort     string        `json:"db_port"`
	DBName     string        `json:"db_name"`
	DBUser     string        `json:"db_user"`
	DBPassword string        `json:"db_password"`
	JWTSecret  string        `json:"jwt_secret"`
	JWTExpiry  time.Duration `json:"jwt_expiry"`
}

func Load() (*Config, error) {
	jwtExpiryStr := getEnv("JWT_EXPIRY")

	jwtExpiry, err := time.ParseDuration(jwtExpiryStr)
	if err != nil {
		return nil, errors.New("invalid JWT_EXPIRY format")
	}

	cfg := &Config{
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
		DBName:     getEnv("DB_NAME"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		JWTSecret:  getEnv("JWT_SECRET"),
		JWTExpiry:  jwtExpiry,
	}

	if cfg.DBHost == "" ||
		cfg.DBPort == "" ||
		cfg.DBName == "" ||
		cfg.DBUser == "" ||
		cfg.DBPassword == "" ||
		cfg.JWTSecret == "" ||
		jwtExpiryStr == "" {
		return nil, errors.New("missing required environment variables")
	}

	return cfg, nil
}

func getEnv(key string) string {
	return os.Getenv(key)
}
