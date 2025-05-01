package models

type UserSessionsData struct {
	SessionId   string   `json:"session_id" bson:"session_id"`
	CodeBaseIds []string `json:"codebase_ids" bson:"codebase_ids"`
}
type User struct {
	UserID       string             `json:"userid" bson:"userid"`
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password" `
	CodebaseData []CodeBaseModel    `json:"codebasedata" bson:"codebasedata"`
	UserSessions []UserSessionsData `json:"user_sessions" bson:"user_sessions"`
}
