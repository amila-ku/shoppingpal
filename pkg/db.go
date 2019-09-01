package shoppingpal

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type table interface {
	put()
	get()
}

type db struct {
	region    string
	tableName string
	endpoint  string
}

func (d db) createTable(w Item) {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String(d.region)}, &aws.Config{Endpoint: aws.String(d.endpoint)})
	table := db.Table(d.tableName)
	err := table.Put(w).Run()

	if err != nil {
		log.Fatal("Failed to write to DB", err)
	}

}
