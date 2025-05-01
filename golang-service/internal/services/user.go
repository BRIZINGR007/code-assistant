package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

func SaveUser(user *models.User) error {
	repo := repositories.GetUserRepository()
	err := repo.AddUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
func GetUserByEmail(email string) (*models.User, error) {
	repo := repositories.GetUserRepository()
	user, err := repo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddCodeBase(userId string, codeBaseData *models.CodeBaseModel) error {
	user_repo := repositories.GetUserRepository()
	codebase_repo := repositories.GetCodeBaseRepository()
	err := user_repo.AddCodeBaseToUser(context.Background(), userId, codeBaseData)
	if err != nil {
		return fmt.Errorf("cannot add CodeBaseData to the mongo document")
	}
	err = codebase_repo.InsertCodeBase(context.Background(), *codeBaseData)
	if err != nil {
		return fmt.Errorf("cannot add CodeBaseData to the codeBase Repository")
	}
	return nil
}

func GetUserCodeBases(userId string) ([]models.CodeBaseModel, error) {
	repo := repositories.GetUserRepository()
	userinfo, err := repo.GetUserByUserId(context.Background(), userId)
	if err != nil {
		return nil, fmt.Errorf("error  in retrieving  code-base data from  User Repo")
	}

	return userinfo.CodebaseData, nil
}

func SyncCodeBases(userId string) error {
	userRepo := repositories.GetUserRepository()
	codebaseRepo := repositories.GetCodeBaseRepository()
	codebases, err := codebaseRepo.GetAllCodeBases(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving code-base data from CodeBase repository: %v", err)
	}

	userInfo, err := userRepo.GetUserByUserId(context.Background(), userId)
	if err != nil {
		return fmt.Errorf("error retrieving user data: %v", err)
	}
	userCodeBases := userInfo.CodebaseData

	// Create a map for faster lookup
	globalCodeBaseMap := make(map[string]models.CodeBaseModel)
	for _, cb := range codebases {
		globalCodeBaseMap[cb.CodeBaseId] = cb
	}

	userCodeBaseMap := make(map[string]models.CodeBaseModel)
	for _, ucb := range userCodeBases {
		userCodeBaseMap[ucb.CodeBaseId] = ucb
	}

	// Add new codebases to user
	for id, cb := range globalCodeBaseMap {
		if _, exists := userCodeBaseMap[id]; !exists {
			err := userRepo.AddCodeBaseToUser(context.Background(), userId, &cb)
			if err != nil {
				return fmt.Errorf("cannot add CodeBaseData to the mongo document: %v", err)
			}
		}
	}

	// Remove codebases that no longer exist in the global list
	for id := range userCodeBaseMap {
		if _, exists := globalCodeBaseMap[id]; !exists {
			err := userRepo.DeleteCodeBaseFromUser(context.Background(), userId, id)
			if err != nil {
				return fmt.Errorf("cannot delete CodeBaseData from the mongo document: %v", err)
			}
		}
	}

	return nil
}

func DeleteCodeBase(userId string, codeBaseId string) error {
	user_repo := repositories.GetUserRepository()
	codecontext_repo := repositories.GetCodeContextRepository()
	codebase_repo := repositories.GetCodeBaseRepository()
	var err error
	err = user_repo.DeleteCodeBaseFromUser(context.Background(), userId, codeBaseId)
	if err != nil {
		return fmt.Errorf("error in deleting the codedb context from user collection")
	}

	codebase_metadata, err := codebase_repo.GetCodeBaseById(context.Background(), codeBaseId)
	if err != nil {
		log.Printf("Error retrieving codebase with ID '%s': %v", codeBaseId, err)
		return nil
	}
	if codebase_metadata.UserId == userId {
		err := codebase_repo.DeleteCodeBaseById(context.Background(), codeBaseId)
		if err != nil {
			return fmt.Errorf("error in deleting Code Base")
		}
		err = codecontext_repo.DeleteCodeContextByCodeBaseId(context.Background(), codeBaseId)
		if err != nil {
			return fmt.Errorf("error in deleting Code Context")
		}

	}
	return nil
}

func AddUserSession(userId string, codeBaseIds []string, chatId string) error {
	userRepo := repositories.GetUserRepository()
	err := userRepo.AddUserSession(context.Background(), userId, codeBaseIds, chatId)
	if err != nil {
		return fmt.Errorf("error in adding chat to user: %w", err)
	}
	return nil
}

func DeleteUserSession(sessionId string, userId string) error {
	userRepo := repositories.GetUserRepository()
	chat_repo := repositories.GetChatRepository()
	err := userRepo.DeleteUserSession(context.Background(), sessionId, userId)
	if err != nil {
		return fmt.Errorf("error in removing chat from user metadata: %w", err)
	}
	err = chat_repo.DeleteChatsBySession(context.Background(), sessionId)
	if err != nil {
		return fmt.Errorf("error in removing chat from chat collection: %w", err)
	}
	return nil
}

func GetAllUserSessions(userId string) ([]models.UserSessionsData, error) {
	repo := repositories.GetUserRepository()
	cd, err := repo.GetUserByUserId(context.Background(), userId)
	if err != nil {
		return nil, fmt.Errorf("error  in retrieving  code-base data")
	}
	return cd.UserSessions, nil
}

func GetSessionCodeBases(sessionId string, userId string) ([]string, error) {
	userRepo := repositories.GetUserRepository()
	userInfo, err := userRepo.GetUserByUserId(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	var sessionCodebaseIds []string
	for _, session := range userInfo.UserSessions {
		if session.SessionId == sessionId {
			sessionCodebaseIds = session.CodeBaseIds
			break
		}
	}
	if len(sessionCodebaseIds) == 0 {
		return nil, errors.New("no codebase ids found for the given session id")
	}

	return sessionCodebaseIds, nil

}
