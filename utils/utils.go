package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/kennybebel/testgostruc/model"
)

func Unmarshal(fileName string, data interface{}) error {

	fileBody, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileBody, &data)
	return err

}

//BatchQuery will query a table for all the items that have the supplierID passed
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
	if err != nil {
		return result, err
	}
	return result, nil
}

//BatchDeleteOrder will delete all orders from the table given
func BatchDeleteOrder(orders []model.OrderInfo, tableName string, client *dynamodb.DynamoDB) error {
	table := make(map[string][]*dynamodb.WriteRequest)
	var orderKey []*dynamodb.WriteRequest
	// looping through the order slice and creating key for each Order to be deleted
	for _, order := range orders {
		item := &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: map[string]*dynamodb.AttributeValue{
					"SupplierID": {
						S: aws.String(order.SupplierID),
					},
					"SortKey": {
						S: aws.String(order.SortKey),
					},
				},
			},
		}
		// adding the Key to items array to use for the BatchWriteItemInput(batch delete)
		orderKey = append(orderKey, item)
	}

	// input holds the values(informtion) need for the items needed to be deleted from the TradeOrders table
	table[tableName] = orderKey
	input := &dynamodb.BatchWriteItemInput{
		RequestItems:           table,
		ReturnConsumedCapacity: aws.String("NONE"),
	}

	// result will hold the results of the items deleted message
	_, err := client.BatchWriteItem(input)
	if err != nil {
		return err
	}
	return err
}
