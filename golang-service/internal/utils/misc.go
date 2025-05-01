package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateHashFromString(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func GenerateUUID() string {
	return uuid.New().String()
}

func AIServiceBaseURL() string {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		log.Fatal("BASE_URL is not set in the environment")
	}
	return baseURL
}

func ReferencesLimit() int {
	limitStr := os.Getenv("REFRENCES_LIMIT")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Println("Invalid REFRENCES_LIMIT, using default %w", 5)
		return 5
	}

	log.Println("REFRENCES_LIMIT loaded as %w", limit)
	return limit
}

func LanguageMap() map[string]string {
	languageMap := map[string]string{
		".py":   "python",
		".js":   "javascript",
		".ts":   "typescript",
		".go":   "go",
		".java": "java",
		".rb":   "ruby",
		".cpp":  "cpp",
		".c":    "c",
		".cs":   "csharp",
		".php":  "php",
		".html": "html",
		".css":  "css",
		".scss": "scss",
		".json": "json",
		".xml":  "xml",
		".sh":   "bash",
		".rs":   "rust",
	}
	return languageMap
}
