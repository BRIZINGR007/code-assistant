package controllers

import (
	"net/http"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/domain"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/services"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/utils"
	"github.com/BRIZINGR007/go-service-utils/helpers"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var input domain.SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Hash pasword ..."})
		return
	}
	user := models.User{
		UserID:       utils.GenerateUUID(),
		Name:         input.Name,
		Email:        input.Email,
		Password:     hashedPassword,
		CodebaseData: []models.CodeBaseModel{},
		UserSessions: []models.UserSessionsData{},
	}
	err = services.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"userid":  user.UserID,
	})

}

func LogIn(c *gin.Context) {
	var input domain.LogInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := services.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get User By Email ..."})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No User Found  by  the Email  ..."})
		return
	}

	passwordIsValid := utils.CheckPasswordHash(input.Password, user.Password)
	if !passwordIsValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is not Matching ..."})
		return
	}

	token, err := helpers.GenerateToken(user.Email, user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}
	setCookie(c, token)
	c.JSON(http.StatusOK, gin.H{"success": "Successfully Logged In ."})

}

func GetAllCodeBases(c *gin.Context) {
	context := helpers.GetGinContextHeadersStruct(c)
	userId := context.UserId
	codeContextData, err := services.GetUserCodeBases(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Cannot Retrieve CodeContext Data."})
	}
	c.JSON(http.StatusOK, codeContextData)
}

func ValidateSession(c *gin.Context) {
	context := helpers.GetGinContextHeadersStruct(c)
	email := context.Email
	userId := context.UserId
	token, err := helpers.GenerateToken(email, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}
	setCookie(c, token)
}

func DeleteCodeBase(c *gin.Context) {
	context := helpers.GetGinContextHeadersStruct(c)
	userId := context.UserId
	codeBaseId := c.Query("codeBaseId")
	err := services.DeleteCodeBase(userId, codeBaseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Servcer  Error."})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Successfully Deleted .."})

}

func CreateUserSession(c *gin.Context) {
	var input domain.ChatToUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context := helpers.GetGinContextHeadersStruct(c)
	userId := context.UserId
	sessionId := utils.GenerateUUID()
	err := services.AddUserSession(userId, input.CodeBaseIds, sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sessionId": sessionId})

}

func DeleteUserSession(c *gin.Context) {
	sessionId := c.Query("sessionId")
	context := helpers.GetGinContextHeadersStruct(c)
	userId := context.UserId
	err := services.DeleteUserSession(sessionId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully Deleted Session"})
}
func GetAllUserSessions(c *gin.Context) {
	context := helpers.GetGinContextHeadersStruct(c)
	usersessions, err := services.GetAllUserSessions(context.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usersessions)
}

func SyncCodeBases(c *gin.Context) {
	context := helpers.GetGinContextHeadersStruct(c)
	err := services.SyncCodeBases(context.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully Synced Databases."})
}

func GetSessionCodeBaseIds(c *gin.Context) {
	sessionId := c.Query("sessionId")
	context := helpers.GetGinContextHeadersStruct(c)
	userId := context.UserId
	codeBaseIds, err := services.GetSessionCodeBases(sessionId, userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": " no CodeBaseIds for  the Session"})
		return
	}
	c.JSON(http.StatusOK, codeBaseIds)
}

func DeleteGeneralChatSession(c *gin.Context) {
	context := helpers.GetGinContextHeadersStruct(c)
	err := services.DeleteChatsBySession(context.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully Deleted General Chats."})
}

func LogOut(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("authorization", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func setCookie(c *gin.Context, token string) {
	// 2 hours = 7200 seconds
	c.SetSameSite(http.SameSiteNoneMode)
	maxAge := 2 * 60 * 60
	c.SetCookie("authorization", token, maxAge, "/", "", true, true)
}
