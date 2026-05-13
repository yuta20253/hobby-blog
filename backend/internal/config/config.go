package config

import "os"

type Config struct {
	UploadPath string
}

func Load() *Config {
	return &Config{
		UploadPath: getEnv("UPLOAD_PATH", "uploads"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
