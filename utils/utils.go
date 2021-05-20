package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
func BatchQueryOrder(tableName string, SupplierID string, filter string, sortkeyName string, client *dynamodb.DynamoDB) (*dynamodb.QueryOutput, error) {
	if filter == "" {
		// Creating the query parameters which will get all test Order from dynamodb
		inputQuery := &dynamodb.QueryInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":SupplierID": {
					S: aws.String(SupplierID),
				},
			},
			KeyConditionExpression: aws.String("SupplierID = :SupplierID"),
			ProjectionExpression:   aws.String("SupplierID, " + sortkeyName),
			TableName:              aws.String(tableName),
		}
		// result will hold the results of the Query of TradeOrders table for all test data
		result, err := client.Query(inputQuery)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	// Creating the query parameters which will get the test Order from dynamodb
	inputQuery := &dynamodb.QueryInput{

		ExpressionAttributeNames: map[string]*string{
			"#Type": aws.String("Type"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":SupplierID": {
				S: aws.String("RAPISAMTEST"),
			},
			":Type": {
				S: aws.String("DELIVERY"),
			},
		},
		KeyConditionExpression: aws.String("SupplierID = :SupplierID"),
		FilterExpression:       aws.String("#Type = :Type"),
		ProjectionExpression:   aws.String("SupplierID, SortKey"),
		TableName:              aws.String(tableName),
	}

	// queryResult will hold the results of the Query of TradeOrders table for all test data
	result, err := client.Query(inputQuery)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Creating the query parameters which will get the test stock from dynamodb


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

func BatchDeleteStock(stocks []model.StockInfo, tableName string, client *dynamodb.DynamoDB) error {
	stockTable := make(map[string][]*dynamodb.WriteRequest)
	var stockItems []*dynamodb.WriteRequest
	// looping through the stocks slice and creating key for each stock to be deleted
	for _, stock := range stocks {
		item := &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: map[string]*dynamodb.AttributeValue{
					"SupplierID": {
						S: aws.String(stock.SupplierID),
					},
					"WarehouseProductID": {
						S: aws.String(stock.WarehouseProductID),
					},
				},
			},
		}
		// adding the Key to items array to use for the BatchWriteItemInput(batch delete)
		stockItems = append(stockItems, item)
	}

	// input holds the values(informtion) need for the items needed to be deleted from the TradeOrders table
	stockTable[tableName] = stockItems
	input := &dynamodb.BatchWriteItemInput{
		RequestItems:           stockTable,
		ReturnConsumedCapacity: aws.String("NONE"),
	}

	// result will hold the results of the items deleted message
	_, err := client.BatchWriteItem(input)
	if err != nil {
		return err
	}
	return err
}

func UploadFileToS3(file *os.File, awsSession *session.Session) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(awsSession)

	// Uploads the file(stored in delivFile) to the S3 bucket.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("rapidtradeinboxs"),
		Key:    aws.String("Deliv.json"),
		Body:   file,
	})
	if err != nil {
		return result, err
	}
	return result, err
}
