package config

import (
	"os"
	"strings"
)

type Config struct {
	UploadPath string
	CORSAllowOrigins []string
	Env string
}

func Load() *Config {
	return &Config{
		UploadPath: getEnv("UPLOAD_PATH", "uploads"),
		CORSAllowOrigins: getEnvAsSlice("CORS_ALLOW_ORIGINS", []string{"http://localhost:5173"})
		Env: getEnv("ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvAsSlice(key string, fallback []string) []string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return strings.Split(val, ",")
}
