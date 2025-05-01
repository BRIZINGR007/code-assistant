package interfcaes

import (
	"context"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

type CodeContextRepositoryInterface interface {
	CodeContextIndexes(ctx context.Context) error
	BulkInsertCodeContext(ctx context.Context, codeContexts *[]models.CodeContext) error
	CodeContextRetriever(ctx context.Context, codeBaseIds []string, embedding []float32, excludeVectorIds []string) ([]models.CodeContext, error)
	DeleteCodeContextByCodeBaseId(ctx context.Context, codeBaseId string) error
}

var _ CodeContextRepositoryInterface = (*repositories.CodeContextRepository)(nil)
