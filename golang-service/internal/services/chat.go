package services

import (
	"context"
	"fmt"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/utils"
)

func AddChat(chat *models.Chat) error {
	repo := repositories.GetChatRepository()

	err := repo.AddChat(context.Background(), chat)
	return err
}

func GetSessionChats(sessionId string) ([]models.Chat, error) {
	repo := repositories.GetChatRepository()
	all_chats, err := repo.GetSessionChats(context.Background(), sessionId)
	if err != nil {
		return nil, err
	}
	return all_chats, nil
}

func DeleteChatsBySession(sessionId string) error {
	repo := repositories.GetChatRepository()
	err := repo.DeleteChatsBySession(context.Background(), sessionId)
	if err != nil {
		return err
	}
	return nil
}

func CheckCodeBaseAvaibility(codeBaseIds []string) error {
	codebaseRepo := repositories.GetCodeBaseRepository()
	existingCodeBases, err := codebaseRepo.GetAllCodeBases(context.Background())
	if err != nil {
		return err
	}
	existingIdsMap := make(map[string]bool)
	for _, codebase := range existingCodeBases {
		existingIdsMap[codebase.CodeBaseId] = true
	}

	for _, id := range codeBaseIds {
		if !existingIdsMap[id] {
			return fmt.Errorf("codebase with ID '%s' not found", id)
		}
	}
	return nil
}

func GeneralSessionChat(user_query string, userId string) (*models.Chat, error) {
	llm_response, err := FetchGeneralLLMResponse(user_query)
	if err != nil {
		return nil, fmt.Errorf("error in Generating LLM Response")
	}
	chatPayload := models.Chat{
		SessionId:    userId,
		ChatId:       utils.GenerateUUID(),
		UserID:       userId,
		AIAnswer:     llm_response,
		UserQuestion: user_query,
	}
	err = AddChat(&chatPayload)
	if err != nil {
		return nil, fmt.Errorf("error in saving chat")
	}
	return &chatPayload, nil

}
