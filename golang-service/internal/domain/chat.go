package domain

type DoSessionContextChat struct {
	SessionId   string   `json:"session_id"`
	CodeBaseIds []string `json:"codebase_ids"`
	Query       string   `json:"query"`
}

type DoGeneralSessionChat struct {
	Query string `json:"query"`
}
