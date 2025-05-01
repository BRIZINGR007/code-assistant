package services

import (
	"context"
	"errors"
	"log"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

func PopulateCodeContext(codeContexts *[]models.CodeContext) error {
	repo := repositories.GetCodeContextRepository()
	err := repo.BulkInsertCodeContext(context.Background(), codeContexts)
	if err != nil {
		return err
	}
	return nil
}

func CodeContextRecursiveRetriever(codeBaseIds []string, query string, accumulated []models.CodeContext, recurseCount int, recurseLimit int) ([]models.CodeContext, error) {
	if recurseCount >= recurseLimit {
		return accumulated, nil
	}
	repo := repositories.GetCodeContextRepository()
	var embedding []float32
	var err error

	if recurseCount == 0 {
		embedding, err = FetchEmbedding(query)
		if err != nil {
			log.Printf("failed to get embeddings for app-002-ai-service")
			return nil, errors.New("failed to get embeddings for app-002-ai-service")
		}
	} else {
		if len(accumulated) == 0 {
			return nil, errors.New("accumulated context is empty in recursive call")
		}
		embedding = accumulated[len(accumulated)-1].Embedding
	}

	var vectors_ids []string
	if len(accumulated) > 0 {
		for _, context := range accumulated {
			vectors_ids = append(vectors_ids, context.VectorId)
		}
	}
	newContexts, err := repo.CodeContextRetriever(context.Background(), codeBaseIds, embedding, vectors_ids)
	if err != nil {
		return nil, err
	}
	rerankedIndex, err := FetchRerankedIndex(query, newContexts)
	if err != nil {
		return nil, err
	}
	accumulated = append(accumulated, newContexts[rerankedIndex])
	recurseCount += 1
	return CodeContextRecursiveRetriever(codeBaseIds, query, accumulated, recurseCount, recurseLimit)
}
