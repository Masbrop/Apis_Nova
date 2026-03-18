package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	AppName         string
	AppEnv          string
	Port            string
	ShutdownTimeout time.Duration
	Database        DatabaseConfig
}

type DatabaseConfig struct {
	Enabled  bool
	Driver   string
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

func Load() (Config, error) {
	cfg := Config{
		AppName:         getEnv("APP_NAME", "apis_nova"),
		AppEnv:          getEnv("APP_ENV", "development"),
		Port:            getEnv("APP_PORT", "8080"),
		ShutdownTimeout: getDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
		Database: DatabaseConfig{
			Enabled:  getBool("DB_ENABLED", false),
			Driver:   getEnv("DB_DRIVER", "postgres"),
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) Validate() error {
	if strings.TrimSpace(c.Port) == "" {
		return fmt.Errorf("APP_PORT is required")
	}

	if !c.Database.Enabled {
		return nil
	}

	if c.Database.Driver != "postgres" {
		return fmt.Errorf("unsupported DB_DRIVER %q", c.Database.Driver)
	}

	required := map[string]string{
		"DB_HOST": c.Database.Host,
		"DB_PORT": c.Database.Port,
		"DB_NAME": c.Database.Name,
		"DB_USER": c.Database.User,
	}

	for key, value := range required {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s is required when DB_ENABLED=true", key)
		}
	}

	return nil
}

func (c Config) HTTPAddress() string {
	return ":" + c.Port
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	return value
}

func getBool(key string, fallback bool) bool {
	value := strings.TrimSpace(strings.ToLower(os.Getenv(key)))
	if value == "" {
		return fallback
	}

	return value == "1" || value == "true" || value == "yes"
}

func getDuration(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return duration
}
