package app

import (
	//"context"
	"time"
	"log"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
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

	// Pretend to do something for 10s
	var result string
	_ = workflow.ExecuteActivity(ctx, DoWork).Get(ctx, &result)
	log.Printf("Got this result from DoWork: " + result)

	// Execute the pause activity
	_ = workflow.ExecuteActivity(ctx, Pause, input).Get(ctx, &result)

	// Resume it by running /start/resume.go

	return result, nil
}
