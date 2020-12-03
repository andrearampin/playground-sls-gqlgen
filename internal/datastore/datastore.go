package datastore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Datastore provides access to the datastore entities.
type Datastore interface {
	CreateRestaurant(name string) (Restaurant, error)
	FetchRestaurant(id string) (Restaurant, error)
	UpdateRestaurant(r Restaurant) (Restaurant, error)
}

// Model interface
type Model interface {
	TableName() string
}

// DB holds the DB connection data.
type DB struct {
	db *dynamodb.DynamoDB
}

// NewDB initiate a new DB connections.
func NewDB() (DB, error) {
	db := dynamodb.New(session.New())
	return DB{db: db}, nil
}

func itemOutputStruct(out dynamodb.GetItemOutput, model Model) error {
	return dynamodbattribute.UnmarshalMap(out.Item, model)
}

func (d DB) putItem(item Model) error {
	marshalled, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = d.db.PutItem(&dynamodb.PutItemInput{
		Item:      marshalled,
		TableName: aws.String(item.TableName()),
	})

	return err
}

func (d DB) getItem(item Model) (*dynamodb.GetItemOutput, error) {
	marshalled, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return nil, err
	}

	result, err := d.db.GetItem(&dynamodb.GetItemInput{
		Key:       marshalled,
		TableName: aws.String(item.TableName()),
	})

	return result, err
}
