package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

func InitMongo() {
	mongoUri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal("MongoDB Connection Error: ", err)
	}

	MongoClient = client
	MongoDatabase = client.Database(dbName)
	init_repositories((MongoDatabase))
	log.Println("Connected to MongoDB ... %w", mongoUri+dbName)
}
func init_repositories(db *mongo.Database) {
	log.Printf("Initializing Repositories ...")
	repositories.InitUserRepository(db)
	repositories.InitCodeContextRepository(db)
	repositories.InitChatRepository(db)
	repositories.InitCodeBaseRepository(db)
}
