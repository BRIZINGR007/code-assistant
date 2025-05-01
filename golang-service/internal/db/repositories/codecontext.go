package repositories

import (
	"context"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CodeContextRepository struct {
	Collection *mongo.Collection
}

func (r *CodeContextRepository) CodeContextIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	codeBaseIdIndex := mongo.IndexModel{
		Keys:    bson.M{"codebase_id": 1},
		Options: options.Index().SetUnique(false),
	}
	vectorIdIndex := mongo.IndexModel{
		Keys:    bson.M{"vector_id": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.Collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		codeBaseIdIndex,
		vectorIdIndex,
	})
	if err != nil {
		return err
	}

	opts := options.SearchIndexes().SetName("vector_index").SetType("vectorSearch")

	type vectorDefinitionField struct {
		Type          string `bson:"type"`
		Path          string `bson:"path"`
		NumDimensions int    `bson:"numDimensions,omitempty"`
		Similarity    string `bson:"similarity,omitempty"`
	}

	type vectorDefinition struct {
		Fields []vectorDefinitionField `bson:"fields"`
	}

	vectorSearchIndexModel := mongo.SearchIndexModel{
		Definition: vectorDefinition{
			Fields: []vectorDefinitionField{
				{
					Type:          "vector",
					Path:          "embedding",
					NumDimensions: 384,
					Similarity:    "dotProduct",
				},
				{
					Type: "filter",
					Path: "codebase_id",
				},
				{
					Type: "filter",
					Path: "vector_id",
				},
			},
		},
		Options: opts,
	}

	_, err = r.Collection.SearchIndexes().CreateOne(ctx, vectorSearchIndexModel)
	if err != nil {
		return err
	}

	return nil
}

func (r *CodeContextRepository) BulkInsertCodeContext(ctx context.Context, codeContexts *[]models.CodeContext) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	batchSize := 1000
	for i := 0; i < len(*codeContexts); i += batchSize {
		end := i + batchSize
		if end > len(*codeContexts) {
			end = len(*codeContexts)
		}
		subBatch := (*codeContexts)[i:end]
		docs := make([]interface{}, len(subBatch))
		for j, cc := range subBatch {
			docs[j] = cc
		}
		_, err := r.Collection.InsertMany(ctx, docs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *CodeContextRepository) CodeContextRetriever(
	ctx context.Context,
	codeBaseIds []string,
	embedding []float32,
	excludeVectorIds []string,
) ([]models.CodeContext, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var codeContextData []models.CodeContext

	filter := bson.M{
		"codebase_id": bson.M{"$in": codeBaseIds},
	}
	if len(excludeVectorIds) > 0 {
		filter["vector_id"] = bson.M{"$nin": excludeVectorIds}
	}

	pipeline := mongo.Pipeline{
		{{Key: "$vectorSearch", Value: bson.D{
			{Key: "index", Value: "vector_index"},
			{Key: "path", Value: "embedding"},
			{Key: "queryVector", Value: embedding},
			{Key: "numCandidates", Value: 75},
			{Key: "limit", Value: 15},
			{Key: "filter", Value: filter},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "vector_id", Value: 1},
			{Key: "hashId", Value: 1},
			{Key: "codebase_id", Value: 1},
			{Key: "codebase_name", Value: 1},
			{Key: "filePath", Value: 1},
			{Key: "code", Value: 1},
			{Key: "embedding", Value: 1},
			{Key: "Score", Value: bson.M{"$meta": "vectorSearchScore"}},
		}}},
	}

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &codeContextData); err != nil {
		return nil, err
	}

	return codeContextData, nil
}

func (r *CodeContextRepository) DeleteCodeContextByCodeBaseId(ctx context.Context, codeBaseId string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{"codebase_id": codeBaseId}
	_, err := r.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
