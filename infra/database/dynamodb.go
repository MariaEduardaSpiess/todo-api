package database

import (
	"fmt"
	"todo-api/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetDynamoDBClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(config.Env.Aws.Endpoint),
		Region:   aws.String(config.Env.Aws.Region),
	}))

	// Create DynamoDB client
	return dynamodb.New(sess)
}

func CreateItem(item any, tableName string) error {
	dynamoItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("got error marshalling item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      dynamoItem,
		TableName: aws.String(tableName),
	}

	dbClient := GetDynamoDBClient()
	_, err = dbClient.PutItem(input)
	if err != nil {
		return fmt.Errorf("got error calling PutItem: %s", err)
	}

	return nil
}

func GetItems(tableName string) ([]map[string]*dynamodb.AttributeValue, error) {
	dbClient := GetDynamoDBClient()
	output, err := dbClient.Scan(&dynamodb.ScanInput{TableName: aws.String(tableName)})
	if err != nil {
		return nil, fmt.Errorf("got error calling Scan: %s", err)
	}

	return output.Items, nil
}
