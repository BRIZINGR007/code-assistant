package models

type CodeBaseModel struct {
	CodeBaseId   string `json:"codebase_id" bson:"codebase_id"`
	CodeBaseName string `json:"codebase_name" bson:"codebase_name"`
	GitHubURL    string `json:"github_url" bson:"github_url"`
	Username     string `json:"username" bson:"username"`
	Branch       string `json:"branch" bson:"branch"`
	FolderPath   string `json:"folder_path" bson:"folder_path"`
	UserId       string `json:"userid" bson:"userid"`
}
