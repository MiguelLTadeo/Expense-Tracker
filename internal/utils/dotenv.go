package utils

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	ex, _ := os.Getwd()

	godotenv.Load(filepath.Join(ex, ".env"), filepath.Join(ex, "../.env"))

	return os.Getenv(key)
}
