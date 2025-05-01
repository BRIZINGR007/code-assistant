package domain

import (
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
)

type RepoRequest struct {
	GitHubURL    string `json:"github_url"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	Branch       string `json:"branch"`
	CodeBaseName string `json:"codebase_name" bson:"codebase_name"`
	FolderPath   string `json:"folder_path"`
}

type FileContent struct {
	FilePath string `json:"filePath"`
	Code     string `json:"code"`
	Chunks   []string
}

type ContextWithSimilarityPayload struct {
	Query   string               `json:"query"`
	Context []models.CodeContext `json:"context"`
}

