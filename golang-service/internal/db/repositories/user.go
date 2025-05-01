package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (r *UserRepository) UserIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.Collection.Indexes().CreateOne(ctx, indexModel)
	return err
}

func (r *UserRepository) AddUser(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, user)
	return err
}
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUserId(ctx context.Context, userId string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"userid": userId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) AddCodeBaseToUser(ctx context.Context, userId string, codeBaseData *models.CodeBaseModel) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"userid": userId, "codebasedata.codeBaseId": bson.M{"$ne": codeBaseData.CodeBaseId}}

	update := bson.M{
		"$addToSet": bson.M{
			"codebasedata": codeBaseData,
		},
	}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		count, err := r.Collection.CountDocuments(ctx, bson.M{"userid": userId})
		if err != nil {
			return err
		}
		if count == 0 {
			return fmt.Errorf("no user found with userId: %s", userId)
		}
		return nil
	}

	return nil
}

func (r *UserRepository) AddUserSession(ctx context.Context, userId string, codeBaseIds []string, sessionId string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"userid": userId}
	sessionData := bson.M{
		"session_id":   sessionId,
		"codebase_ids": codeBaseIds,
	}
	update := bson.M{
		"$push": bson.M{
			"user_sessions": sessionData,
		},
	}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no user found with userId: %s", userId)
	}

	return nil
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, sessionId string, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"userid": userId}
	update := bson.M{
		"$pull": bson.M{
			"user_sessions": bson.M{"session_id": sessionId},
		},
	}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to remove chat_id from user metadata: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteCodeBaseFromUser(ctx context.Context, userId string, codeBaseId string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"userid": userId}
	update := bson.M{
		"$pull": bson.M{
			"codebasedata": bson.M{"codebase_id": codeBaseId},
		},
	}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to remove codebase_id from user metadata: %w", err)
	}
	return nil

}
