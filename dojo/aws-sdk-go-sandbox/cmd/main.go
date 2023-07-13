package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// AWS SDK GO SANDBOX
func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	// create dynamoDB client
	service := dynamodb.NewFromConfig(cfg)

	// Get a list of DynamoDB tables
	response, err := service.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	// List Tables
	fmt.Println("Tables:")
	for _, tableName := range response.TableNames {
		fmt.Println(tableName)
	}
}
