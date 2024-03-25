package main

import (
	"account-service/app"
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "get-account-workflow",
		TaskQueue: app.GetAccountTaskQueue,
	}
	// start of the workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GetAccountWorkflow)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// Get the results
	var Accounts []app.Account
	err = we.Get(context.Background(), &Accounts)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(Accounts, we.GetID(), we.GetRunID())
}

func printResults(accounts []app.Account, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	for _, account := range accounts {
		fmt.Printf("Account Name: %s, Email: %s, Role: %s, Projects: %v, Devices: %v \n", account.Name, account.Email, account.Role, account.Projects, account.Devices)
	}
}
