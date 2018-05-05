package config

import (
	"os"
)

var (
	MongoURL        string
	MongoDB         string
	MongoCollection string
)

func init() {
	MongoURL = os.Getenv("mongo_url")
	MongoDB = os.Getenv("mongo_db")
	MongoCollection = os.Getenv("mongo_collection")
}
