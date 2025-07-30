package app

import (
	"context"
	"time"
	"log"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.temporal.io/sdk/client"
)

func Workflow(ctx workflow.Context, input PaymentDetails) (string, error) {
	log.Printf("WORKFLOW STARTED!")
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        0, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	// Execute the pause activity
	//var result string
	_ = workflow.ExecuteActivity(ctx, Pause, input)//.Get(ctx, &result)

	// Grab a client to complete the activity
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Printf("Error!!!!!!!!!!", err)
	}

	// Pretend to do something for 10s
	var taskToken []byte
	_ = workflow.ExecuteActivity(ctx, DoWork, c).Get(ctx, &taskToken)
	
	// Now that we've done 10s of work, we can finish
	c.CompleteActivity(context.Background(), taskToken, "Complete!", nil)

	return "Complete!", nil
}
