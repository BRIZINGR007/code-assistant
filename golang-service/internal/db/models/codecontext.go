package models

type CodeContext struct {
	VectorId     string    `json:"vector_id" bson:"vector_id"`
	CodeBaseId   string    `json:"codebase_id" bson:"codebase_id"`
	CodeBaseName string    `json:"codebase_name" bson:"codebase_name"`
	HashId       string    `json:"hashId" bson:"hashId"`
	FilePath     string    `json:"filePath" bson:"filePath"`
	Code         string    `json:"code" bson:"code"`
	Embedding    []float32 `json:"embedding" bson:"embedding"`
}
