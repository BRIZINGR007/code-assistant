package repositories

import (
	"context"

	"github.com/BRIZINGR007/go-service-utils/ioc"
	"go.mongodb.org/mongo-driver/mongo"
)

var userRepoSingleton ioc.Singleton[UserRepository]
var codeContextRepoSingleton ioc.Singleton[CodeContextRepository]
var chatRepoSingleton ioc.Singleton[ChatRepository]
var codebaseRepoSingleton ioc.Singleton[CodeBaseRepository]

func InitUserRepository(db *mongo.Database) {
	userRepoSingleton.Get(func() *UserRepository {
		repo := &UserRepository{
			Collection: db.Collection("users"),
		}

		err := repo.UserIndexes(context.Background())
		if err != nil {
			panic("Failed to create index on email: " + err.Error())
		}

		return repo
	})
}

func GetUserRepository() *UserRepository {
	instance := userRepoSingleton.Get(nil)
	if instance == nil {
		panic("UserRepository not initialized. Call InitUserRepository first.")
	}
	return instance
}

func InitCodeContextRepository(db *mongo.Database) {
	codeContextRepoSingleton.Get(func() *CodeContextRepository {
		repo := &CodeContextRepository{
			Collection: db.Collection("codecontexts"),
		}

		err := repo.CodeContextIndexes(context.Background())
		if err != nil {
			panic("Failed to create index on email: " + err.Error())
		}

		return repo
	})
}

func GetCodeContextRepository() *CodeContextRepository {
	instance := codeContextRepoSingleton.Get(nil)
	if instance == nil {
		panic("CodeAssistRepository not initialized. Call InitCodeAssistRepository first.")
	}
	return instance
}

func InitChatRepository(db *mongo.Database) {
	chatRepoSingleton.Get(func() *ChatRepository {
		repo := &ChatRepository{
			Collection: db.Collection("chats"),
		}

		err := repo.ChatIndexes(context.Background())
		if err != nil {
			panic("Failed to create index on email: " + err.Error())
		}

		return repo
	})
}

func GetChatRepository() *ChatRepository {
	instance := chatRepoSingleton.Get(nil)
	if instance == nil {
		panic("ChatRepository not initialized. Call InitCodeAssistRepository first.")
	}
	return instance
}

func InitCodeBaseRepository(db *mongo.Database) {
	codebaseRepoSingleton.Get(func() *CodeBaseRepository {
		repo := &CodeBaseRepository{
			Collection: db.Collection("codebases"),
		}
		err := repo.CodeBaseIndexes(context.Background())
		if err != nil {
			panic("Failed to create index on email: " + err.Error())
		}
		return repo
	})
}

func GetCodeBaseRepository() *CodeBaseRepository {
	instance := codebaseRepoSingleton.Get(nil)
	if instance == nil {
		panic("Codebase Repository not initialized. Call InitCodeBaseRepository first.")
	}
	return instance
}
