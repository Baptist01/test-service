package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GetAccountWorkflow(ctx workflow.Context) ([]Account, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result []Account
	err := workflow.ExecuteActivity(ctx, CreateMockAccounts).Get(ctx, &result)

	return result, err
}
