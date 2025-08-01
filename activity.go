package app

import (
	"context"
	"log"
	"time"
	"os"

	"go.temporal.io/sdk/activity"
)

func Pause(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Starting the Pause activity")

	taskToken := activity.GetInfo(ctx).TaskToken
	log.Printf("Enjoy your pause! Writing taskToken to file")

	f, err := os.Create("./tmp/tasktoken")
		if err != nil {
		log.Printf("Failed to create file", err)
	}

	defer f.Close()

	_, err = f.Write(taskToken)
	if err != nil {
		log.Printf("Failed to write token to file", err)
	}
	
	f.Sync()

	return "Paused", activity.ErrResultPending
}

func DoWork(ctx context.Context) (string, error) {
	// Wait for 5 seconds to simulate some manual task that needs to be done
	log.Printf("Waiting 5s")
	time.Sleep(5 * time.Second)
	log.Printf("Finished Waiting")

	return "Finished Doing Work", nil
}
