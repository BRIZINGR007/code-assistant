package services

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/domain"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/utils"
	"github.com/remeh/sizedwaitgroup"
)

func clone_and_extract_code(payload *models.CodeBaseModel, token string) ([]domain.FileContent, error) {
	tmpDir, err := os.MkdirTemp("", "repo")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	cloneURL := payload.GitHubURL
	if token != "" && payload.Username != "" {
		parts := strings.Split(payload.GitHubURL, "https://")
		if len(parts) > 1 {
			cloneURL = fmt.Sprintf("https://%s:%s@%s", payload.Username, token, parts[1])
		}
	}
	var args []string
	if payload.Branch != "" {
		args = []string{"clone", "-b", payload.Branch, "--single-branch", cloneURL, tmpDir}
	} else {
		args = []string{"clone", cloneURL, tmpDir}
	}

	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("git clone failed: %w, output: %s", err, output)
	}

	basePath := tmpDir
	if payload.FolderPath != "" {
		basePath = filepath.Join(tmpDir, payload.FolderPath)
		if _, err := os.Stat(basePath); os.IsNotExist(err) {
			return nil, fmt.Errorf("specified folder path does not exist: %s", payload.FolderPath)
		}
	}

	var files []domain.FileContent

	err = filepath.Walk(basePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(basePath, path)
			if err != nil {
				fmt.Printf("Error getting relative path for %s: %v\n", path, err)
				return nil
			}
			subString := "_init_.py"
			if strings.Contains(relPath, subString) {
				return nil
			}

			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", path, err)
				return nil
			}
			files = append(files, domain.FileContent{
				FilePath: relPath,
				Code:     string(content),
			})
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}
	log.Println("FILES  ....", files)
	return files, nil
}

func processCodeChunk(filePath, chunk string, lang string, hasLang bool, codeBaseId string, codeBaseName string) (models.CodeContext, error) {
	var formattedChunk string
	if hasLang {
		formattedChunk = fmt.Sprintf("```%s\n%s\n```", lang, chunk)
	} else {
		formattedChunk = fmt.Sprintf("```\n%s\n```", chunk)
	}

	embedding, err := FetchEmbedding(formattedChunk)
	if err != nil {
		return models.CodeContext{}, err
	}

	context := models.CodeContext{
		VectorId:     utils.GenerateUUID(),
		CodeBaseId:   codeBaseId,
		CodeBaseName: codeBaseName,
		HashId:       utils.GenerateHashFromString(chunk),
		FilePath:     filePath,
		Code:         formattedChunk,
		Embedding:    embedding,
	}

	return context, nil
}

func parallelProcessChunks(tasks []func() (models.CodeContext, error)) ([]models.CodeContext, error) {
	swg := sizedwaitgroup.New(30)
	resultChan := make(chan models.CodeContext)
	errChan := make(chan error, 1)

	go func() {
		swg.Wait()
		close(resultChan)
	}()

	for _, task := range tasks {
		swg.Add()
		go func(t func() (models.CodeContext, error)) {
			defer swg.Done()

			result, err := t()
			if err != nil {
				select {
				case errChan <- err:
				default:
				}
				return
			}
			select {
			case resultChan <- result:
			case <-errChan: // error already occurred, bail
			}
		}(task)
	}

	var contexts []models.CodeContext
	var firstError error

	for result := range resultChan {
		contexts = append(contexts, result)
	}

	select {
	case firstError = <-errChan:
		return nil, firstError
	default:
		return contexts, nil
	}
}

func generate_code_contexts(filecontents *[]domain.FileContent, codeBaseId string, codeBaseName string) ([]models.CodeContext, error) {
	languageMap := utils.LanguageMap()
	var tasks []func() (models.CodeContext, error)

	for _, file := range *filecontents {
		ext := filepath.Ext(file.FilePath)
		lang, hasLang := languageMap[ext]

		for _, chunk := range file.Chunks {
			filePath, chunkText, langVal, hasLangVal := file.FilePath, chunk, lang, hasLang
			task := func() (models.CodeContext, error) {
				return processCodeChunk(filePath, chunkText, langVal, hasLangVal, codeBaseId, codeBaseName)
			}

			tasks = append(tasks, task)
		}
	}
	// Process all tasks with a maximum concurrency of 20
	return parallelProcessChunks(tasks)
}

func split_code_into_chunks(input string) []string {
	re := regexp.MustCompile(`[^\S\r\n]+`)
	input = re.ReplaceAllString(input, " ")
	lines := strings.Split(input, "\n")
	var result []string
	var currentChunk strings.Builder
	wordCount := 0

	for _, line := range lines {
		words := strings.Fields(line)
		lineWordCount := len(words)
		if wordCount+lineWordCount > 300 && wordCount > 0 {
			result = append(result, currentChunk.String())
			currentChunk.Reset()
			wordCount = 0
		}

		if currentChunk.Len() > 0 {
			currentChunk.WriteString("\n")
		}
		currentChunk.WriteString(line)
		wordCount += lineWordCount
	}
	if currentChunk.Len() > 0 {
		result = append(result, currentChunk.String())
	}

	return result

}

func ProcessCodebaseForEmbedding(payload *models.CodeBaseModel, token string) ([]models.CodeContext, error) {
	fileContents, err := clone_and_extract_code(payload, token)
	if err != nil {
		return nil, fmt.Errorf("error  In Extracting  Code : %v", err)
	}
	processedContents := make([]domain.FileContent, len(fileContents))
	for i, content := range fileContents {
		processedContent := content
		chunks := split_code_into_chunks(content.Code)
		processedContent.Chunks = chunks
		processedContent.FilePath = payload.CodeBaseName + "/" + content.FilePath
		processedContents[i] = processedContent
	}
	code_context_with_embeddings, err := generate_code_contexts(&processedContents, payload.CodeBaseId, payload.CodeBaseName)
	if err != nil {
		return nil, fmt.Errorf("error  in populating context with embeddings : %v", err.Error())
	}
	return code_context_with_embeddings, nil
}
