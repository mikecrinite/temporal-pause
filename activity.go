package app

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/activity"
)

// This should probably be stored outside the service. This is just an example
var TokenMap = map[string][]byte{}

func Pause(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Starting the Pause activity")

	// Simulate "things" happening
	time.Sleep(1 * time.Second)

	taskToken := activity.GetInfo(ctx).TaskToken
	log.Printf("Enjoy your pause!")
	TokenMap["taskToken"] = taskToken

	return "", activity.ErrResultPending
}

func DoWork(ctx context.Context) ([]byte, error) {
	// Wait for 10 seconds to simulate some manual task that needs to be done
	log.Printf("Waiting 10s")
	time.Sleep(10 * time.Second)
	log.Printf("Finished Waiting")

	// Complete activity
	taskToken := TokenMap["taskToken"]
	log.Printf("Retrieved this token: " + string(taskToken))

	return taskToken, nil
}
