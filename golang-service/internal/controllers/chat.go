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

func FetchSessionChats(c *gin.Context) {
	sessionId := c.Query("sessionId")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required codeBaseId parameter"})
		return
	}

	allChats, err := services.GetSessionChats(sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve chats", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allChats)
}

func FetchGeneralSessionChats(c *gin.Context) {
	context := helpers.GetGinContextHeadersStruct(c)
	allChats, err := services.GetSessionChats(context.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve chats", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allChats)
}

func GeneralSessionChat(c *gin.Context) {
	var input domain.DoGeneralSessionChat
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid  Request."})
		return
	}
	context := helpers.GetGinContextHeadersStruct(c)
	chatPayload, err := services.GeneralSessionChat(input.Query, context.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chatPayload)
}

func DoSessionContextChat(c *gin.Context) {
	var input domain.DoSessionContextChat
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Format"})
		return
	}
	err := services.CheckCodeBaseAvaibility(input.CodeBaseIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CodeBase Id has  been deleted . Please sync codebases and cerate  a  new session ."})
		return
	}
	referenceslimit := utils.ReferencesLimit()

	codeContext, err := services.CodeContextRecursiveRetriever(input.CodeBaseIds, input.Query, nil, 0, referenceslimit)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	codeContextWithSimilarity, err := services.FetchContextWithSimilarity(&codeContext, input.Query)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	answer, err := services.FetchContextLLMResponse(input.Query, codeContext)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Retrieve LLM Answer ."})
		return
	}
	context := helpers.GetGinContextHeadersStruct(c)
	chatPayload := models.Chat{
		SessionId:    input.SessionId,
		ChatId:       utils.GenerateUUID(),
		UserID:       context.UserId,
		AIAnswer:     answer,
		UserQuestion: input.Query,
		References:   codeContextWithSimilarity,
	}
	err = services.AddChat(&chatPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save chat", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chatPayload)

}
