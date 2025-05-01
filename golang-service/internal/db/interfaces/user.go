package interfcaes

import (
	"context"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

type UserRepositoryInterface interface {
	UserIndexes(ctx context.Context) error
	AddUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUserId(ctx context.Context, userId string) (*models.User, error)
	AddCodeBaseToUser(ctx context.Context, userId string, codeBaseData *models.CodeBaseModel) error
	DeleteCodeBaseFromUser(ctx context.Context, userId string, codeBaseId string) error
	AddUserSession(ctx context.Context, userId string, codeBaseIds []string, sessionId string) error
	DeleteUserSession(ctx context.Context, sessionId string, userId string) error
}

var _ UserRepositoryInterface = (*repositories.UserRepository)(nil)
