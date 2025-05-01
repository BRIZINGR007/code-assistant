package main

import (
	"log"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Bootstrap  background workers and  other  services
	bootstrap()
	server := gin.Default()
	//Add  CORS  middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4280"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowAllOrigins:  true,
	}))
	routes.SetupRoutes(server)
	// Start Server
	if err := server.Run(":4281"); err != nil {
		log.Fatalf("Failed to  start server  : %v", err)
	}

}
