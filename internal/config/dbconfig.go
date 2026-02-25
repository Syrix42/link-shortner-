package config

import (
	"fmt"
	"os"

	"github.com/Syrix42/link-shortener/internal/infra/database"
	"github.com/joho/godotenv"
)

type DBEnv struct {
	EnvFile string
}

func LoadDBConfig(envFile string) (database.Config, error) {
	if envFile != "" {
		_ = godotenv.Load(envFile)

	} else {
		_ = godotenv.Load()
	}
	cfg := database.Config{
		Host:     getEnv("DB_HOST", ""),
		Port:     getEnv("DB_PORT", "5432"),
		Name:     getEnv("DB_NAME", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
		TimeZone: getEnv("DB_TIMEZONE", "UTC"),
	}
	missing := missingVars(map[string]string{
		"DB_HOST":     cfg.Host,
		"DB_NAME":     cfg.Name,
		"DB_USER":     cfg.User,
		"DB_PASSWORD": cfg.Password,
	})
	if len(missing) > 0 {
		return database.Config{}, fmt.Errorf("missing required env vars: %v", missing)
	}

	return cfg, nil

}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func missingVars(vars map[string]string) []string {
	var missing []string
	for k, v := range vars {
		if v == "" {
			missing = append(missing, k)
		}
	}
	return missing
}
