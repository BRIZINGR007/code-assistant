package interfcaes

import (
	"context"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

type ChatRepositoryInterface interface {
	ChatIndexes(ctx context.Context) error
	AddChat(ctx context.Context, chat *models.Chat) error
	GetSessionChats(ctx context.Context, sessionId string) ([]models.Chat, error)
	DeleteChatsBySession(ctx context.Context, sessionId string) error
}

var _ ChatRepositoryInterface = (*repositories.ChatRepository)(nil)
