package domain

type Context struct {
	Code     string `json:"Code"`
	FilePath string `json:"FilePath"`
}
type LLMRequest struct {
	Query    string    `json:"query"`
	Contexts []Context `json:"contexts"`
}

type RerankingPayload struct {
	Query    string   `json:"query"`
	Contexts []string `json:"contexts"`
}
