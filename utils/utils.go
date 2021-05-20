package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Unmarshal(fileName string, data interface{}) error {

	fileBody, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileBody, &data)
	return err

}

func BatchQuery(tableName string, SupplierID string, client *dynamodb.DynamoDB) (*dynamodb.QueryOutput, error) {
	inputQuery := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":SupplierID": {
				S: aws.String(SupplierID),
			},
		},
		KeyConditionExpression: aws.String("SupplierID = :SupplierID"),
		ProjectionExpression:   aws.String("SupplierID, SortKey"),
		TableName:              aws.String(tableName),
	}
	// result will hold the results of the Query of TradeOrders table for all test data
	result, err := client.Query(inputQuery)
	//log the query
	if err != nil {
		return result, err
	}
	return result, nil
}
