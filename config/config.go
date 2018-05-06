package config

import (
	"os"
)

var (
	MongoURL        string
	MongoDB         string
	MongoCollection string
	NotifyAddress   string
)

func init() {
	MongoURL = os.Getenv("MONGO_URL")
	MongoDB = os.Getenv("MONGO_DB")
	MongoCollection = os.Getenv("MONGO_COLLECTION")
	NotifyAddress = os.Getenv("EMAIL_ADDRESS")
}
