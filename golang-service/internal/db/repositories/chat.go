package repositories

import (
	"context"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	Collection *mongo.Collection
}

func (r *ChatRepository) ChatIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sessionIdIndex := mongo.IndexModel{
		Keys: bson.M{"session_id": 1},
	}

	_, err := r.Collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		sessionIdIndex,
	})
	return err
}

func (r *ChatRepository) AddChat(ctx context.Context, chat *models.Chat) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.Collection.InsertOne(ctx, chat)
	return err
}

func (r *ChatRepository) GetSessionChats(ctx context.Context, sessionId string) ([]models.Chat, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{
		"session_id": sessionId,
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var chats []models.Chat
	if err := cursor.All(ctx, &chats); err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *ChatRepository) DeleteChatsBySession(ctx context.Context, sessionId string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"session_id": sessionId}

	_, err := r.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
