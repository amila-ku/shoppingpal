package shoppingpal

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"

    "fmt"
    "os"
)

type table interface {
	put()
	get()
}

type db struct {
	region string,
	tableName string

}

type dynamodbtable struct {
	name string
}

func (t dynamodbtable) put(w Items) {
	err := t.Put(w).Run() 
}


func (d db) createTable() {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String(d.region)})
	table := db.Table(d.tableName)

	return table
	
}