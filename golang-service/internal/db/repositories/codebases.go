package repositories

import (
	"context"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CodeBaseRepository struct {
	Collection *mongo.Collection
}

func (r *CodeBaseRepository) CodeBaseIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sessionIdIndex := mongo.IndexModel{
		Keys:    bson.M{"codebase_id": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.Collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		sessionIdIndex,
	})
	return err
}

func (r *CodeBaseRepository) GetAllCodeBases(ctx context.Context) ([]models.CodeBaseModel, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var codebases []models.CodeBaseModel
	if err := cursor.All(ctx, &codebases); err != nil {
		return nil, err
	}

	return codebases, nil
}

func (r *CodeBaseRepository) InsertCodeBase(ctx context.Context, codebase models.CodeBaseModel) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, codebase)
	return err
}

func (r *CodeBaseRepository) DeleteCodeBaseById(ctx context.Context, codebaseId string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"codebase_id": codebaseId}

	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *CodeBaseRepository) GetCodeBaseById(ctx context.Context, codebaseId string) (*models.CodeBaseModel, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"codebase_id": codebaseId}

	var codebase models.CodeBaseModel
	err := r.Collection.FindOne(ctx, filter).Decode(&codebase)
	if err != nil {
		return nil, err
	}

	return &codebase, nil
}
