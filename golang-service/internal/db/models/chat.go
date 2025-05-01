package models

type ReferencesWithSimilarity struct {
	VectorId              string    `json:"vector_id" bson:"vector_id"`
	CodeBaseId            string    `json:"codebase_id" bson:"codebase_id"`
	CodeBaseName          string    `json:"codebase_name" bson:"codebase_name"`
	HashId                string    `json:"hashId" bson:"hashId"`
	FilePath              string    `json:"filePath" bson:"filePath"`
	Code                  string    `json:"code" bson:"code"`
	Embedding             []float32 `json:"embedding" bson:"embedding"`
	SimilarityToQuery     float32   `json:"similarity_to_query" bson:"similarity_to_query"`
	SimilarityToPrevQuery float32   `json:"similarity_to_prev_query" bson:"similarity_to_prev_query"`
}

type Chat struct {
	SessionId    string                     `json:"session_id" bson:"session_id"`
	ChatId       string                     `json:"chat_id" bson:"chat_id"`
	UserID       string                     `json:"user_id" bson:"user_id"`
	AIAnswer     string                     `json:"ai_answer" bson:"ai_answer"`
	UserQuestion string                     `json:"user_question" bson:"user_question"`
	References   []ReferencesWithSimilarity `json:"references" bson:"refrences"`
}
