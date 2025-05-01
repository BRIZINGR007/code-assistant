package routes

import (
	"github.com/BRIZINGR007/app-002-code-assistant/internal/controllers"
	"github.com/BRIZINGR007/go-service-utils/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	baseServer := server.Group("/code-assistant")
	baseServer.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "OK"})
	})
	//User Routes
	userRoutes(baseServer)
}

func userRoutes(baseServer *gin.RouterGroup) {
	users_base_path := "/users"

	//Un-Authenticated Routes  ...
	user_ua := baseServer.Group(users_base_path)
	user_ua.POST("/signup", controllers.SignUp)
	user_ua.POST("/login", controllers.LogIn)
	user_ua.POST("/logout", controllers.LogOut)

	user := baseServer.Group(users_base_path)
	user.Use(middlewares.RestMiddleware)
	user.GET("/validate-session", controllers.ValidateSession)
	user.GET("/get-code-bases", controllers.GetAllCodeBases)
	user.GET("/get-session-codebase-ids", controllers.GetSessionCodeBaseIds)
	user.DELETE("/delete-codebase", controllers.DeleteCodeBase)
	user.POST("/create-session", controllers.CreateUserSession)
	user.DELETE("/delete-chat-session", controllers.DeleteUserSession)
	user.DELETE("/delete-general-chat-session", controllers.DeleteGeneralChatSession)
	user.GET("/get-chat-sessions", controllers.GetAllUserSessions)
	user.POST("/sync-codebases", controllers.SyncCodeBases)

	//Authenticated Routes  ...
	codeassist_base_path := "/code-assist"
	codeassist := baseServer.Group(codeassist_base_path)
	codeassist.Use(middlewares.RestMiddleware)
	codeassist.POST("/extract-code", controllers.GitCodeExtractor)
	codeassist.POST("/session-chat", controllers.DoSessionContextChat)
	codeassist.POST("/general-session-chat", controllers.GeneralSessionChat)

	//AuthenticatedRoutes
	chat_base_path := "/chat"
	chat := baseServer.Group(chat_base_path)
	chat.Use(middlewares.RestMiddleware)
	chat.GET("get-session-chats", controllers.FetchSessionChats)
	chat.GET("get-general-session-chats", controllers.FetchGeneralSessionChats)

}
