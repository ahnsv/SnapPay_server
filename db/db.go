package db

import (
	"github.com/ahnsv/snappay-server/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamodb.DynamoDB

// DB Init
func Init() {
	c := config.GetConfig()
	db = dynamodb.New(session.New(&aws.Config{
		Region:      aws.String(c.GetString("db.region")),
		Credentials: credentials.NewStaticCredentials(c.GetString("db.keyId"), c.GetString("db.secret"), "snappay-server"),
	}))
}

// grab db
func GetDB() *dynamodb.DynamoDB {
	Init()
	return db
}
