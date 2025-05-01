package main

import (
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db"
)

func bootstrap() {
	//SQS Polling
	// queue.InitEventProcessor()
	// queueURL := os.Getenv("SQS_QUEUE_URL")
	// if queueURL == "" {
	// 	log.Fatal("SQS_QUEUE_URL not found in environment variables")
	// }
	// processor := &sqs_client.SQSProcessor{QueueURL: queueURL}
	// go processor.StartPolling()

	//Mongo
	db.InitMongo()
}
