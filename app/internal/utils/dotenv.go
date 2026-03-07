package utils

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	ex, _ := os.Getwd()
	// Tenta carregar o .env da pasta atual ou de uma acima
	godotenv.Load(filepath.Join(ex, ".env"), filepath.Join(ex, "../.env"))

	return os.Getenv(key)
}
