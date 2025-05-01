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

func GitCodeExtractor(c *gin.Context) {
	var req domain.RepoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Format"})
		return
	}
	context := helpers.GetGinContextHeadersStruct(c)
	codeBaseDataPayload := &models.CodeBaseModel{
		CodeBaseId:   utils.GenerateUUID(),
		GitHubURL:    req.GitHubURL,
		Username:     req.Username,
		CodeBaseName: req.CodeBaseName,
		Branch:       req.Branch,
		FolderPath:   req.FolderPath,
		UserId:       context.UserId,
	}
	codeContextWithEmbeddings, err := services.ProcessCodebaseForEmbedding(codeBaseDataPayload, req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in  populating  the code  context"})
		return
	}
	if err := services.PopulateCodeContext(&codeContextWithEmbeddings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to  populate context . %v"})
		return
	}
	err = services.AddCodeBase(context.UserId, codeBaseDataPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user does  not exist ."})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": "Proccessed ."})
}
