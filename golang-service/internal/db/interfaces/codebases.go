package interfcaes

import (
	"context"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

type CodeBaseRepositoryInterface interface {
	CodeBaseIndexes(ctx context.Context) error
	GetAllCodeBases(ctx context.Context) ([]models.CodeBaseModel, error)
	InsertCodeBase(ctx context.Context, codebase models.CodeBaseModel) error
	DeleteCodeBaseById(ctx context.Context, codebaseId string) error
	GetCodeBaseById(ctx context.Context, codebaseId string) (*models.CodeBaseModel, error)
}

var _ CodeBaseRepositoryInterface = (*repositories.CodeBaseRepository)(nil)
