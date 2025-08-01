package main

import (
	"context"
	"log"
	"os"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	taskTokenFileName := "./tmp/tasktoken"
	taskToken, err := os.ReadFile(taskTokenFileName)
	if err != nil {
		log.Printf("Could not read file")
	} else {
		log.Printf("TaskToken read: " + string(taskToken))
	}

	err = os.Remove(taskTokenFileName)

	c.CompleteActivity(context.Background(), taskToken, "Complete!", nil)
}

