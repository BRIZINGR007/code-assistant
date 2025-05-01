package services

import (
	"encoding/json"
	"fmt"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/domain"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/utils"
)

func FetchContextLLMResponse(userQuery string, codeContexts []models.CodeContext) (string, error) {
	url := utils.AIServiceBaseURL() + "/api/app002-ai-service/decoder/get-llmresponse"
	contexts := make([]domain.Context, len(codeContexts))
	for i, ctx := range codeContexts {
		contexts[i] = domain.Context{
			Code:     ctx.Code,
			FilePath: ctx.FilePath,
		}
	}
	reqBody := domain.LLMRequest{
		Query:    userQuery,
		Contexts: contexts,
	}
	respBody, err := utils.HttpRequest("POST", url, nil, reqBody, nil)
	if err != nil {
		return "", err
	}

	var llmresponse string
	if err := json.Unmarshal(respBody, &llmresponse); err != nil {
		return "", fmt.Errorf("error parsing JSON: %w", err)
	}
	return llmresponse, nil
}

func FetchGeneralLLMResponse(userQuery string) (string, error) {
	url := utils.AIServiceBaseURL() + "/api/app002-ai-service/decoder/general-chat"
	queryParams := map[string]string{
		"user_query": userQuery,
	}
	respBody, err := utils.HttpRequest("GET", url, queryParams, nil, nil)
	if err != nil {
		return "", err
	}
	var llmresponse string
	if err := json.Unmarshal(respBody, &llmresponse); err != nil {
		return "", fmt.Errorf("error parsing JSON: %w", err)
	}
	return llmresponse, nil
}

func FetchEmbedding(query string) ([]float32, error) {
	url := utils.AIServiceBaseURL() + "/api/app002-ai-service/encoder/get-embedding"

	queryParams := map[string]string{
		"query": query,
	}
	responseBody, err := utils.HttpRequest("GET", url, queryParams, nil, nil)
	if err != nil {
		return nil, err
	}
	var embedding []float32
	if err := json.Unmarshal(responseBody, &embedding); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	return embedding, nil
}

func FetchRerankedIndex(query string, retrievedContexts []models.CodeContext) (int, error) {

	url := utils.AIServiceBaseURL() + "/api/app002-ai-service/cross-encoder/reranking"
	contexts := make([]string, len(retrievedContexts))
	for i, ctx := range retrievedContexts {
		contexts[i] = "filePath :" + ctx.FilePath + "\n\n" + ctx.Code
	}
	payload := domain.RerankingPayload{
		Query:    query,
		Contexts: contexts,
	}

	responseBody, err := utils.HttpRequest("POST", url, nil, payload, nil)
	if err != nil {
		return -1, err
	}

	var maxIndex int
	if err := json.Unmarshal(responseBody, &maxIndex); err != nil {
		return -1, fmt.Errorf("error parsing JSON: %w", err)
	}

	return maxIndex, nil
}

func FetchContextWithSimilarity(contexts *[]models.CodeContext, query string) ([]models.ReferencesWithSimilarity, error) {
	url := utils.AIServiceBaseURL() + "/api/app002-ai-service/encoder/get-cosine-similarity-scores"
	requestBody := domain.ContextWithSimilarityPayload{
		Query:   query,
		Context: *contexts,
	}
	respBody, err := utils.HttpRequest("POST", url, nil, requestBody, nil)
	if err != nil {
		return nil, err
	}
	var result []models.ReferencesWithSimilarity
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return result, nil
}
