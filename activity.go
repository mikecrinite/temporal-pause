package app

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/activity"
)

var TokenMap = map[string][]byte{}

func Pause(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Starting an activity and pausing it immediately")

	taskToken := activity.GetInfo(ctx).TaskToken
	log.Printf("Got a task token!")
	TokenMap["taskToken"] = taskToken

	return "", activity.ErrResultPending
}

func DoWork(ctx context.Context) ([]byte, error) {
	// Wait
	log.Printf("Waiting 10s")
	time.Sleep(10 * time.Second)
	log.Printf("Finished Waiting")

	// Complete activity
	taskToken := TokenMap["taskToken"]
	log.Printf("Retrieved this token: " + string(taskToken))

	return taskToken, nil
}
