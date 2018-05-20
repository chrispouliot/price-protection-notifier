package config

import (
	"os"
)

var (
	MongoURL          string
	MongoDB           string
	MongoCollection   string
	MailNotifyAddress string
	MailApiUrl        string
	MailAPIKey        string
	MailFromAddress   string
)

func Init() {
	MongoURL = os.Getenv("MONGO_URL")
	MongoDB = os.Getenv("MONGO_DB")
	MongoCollection = os.Getenv("MONGO_COLLECTION")
	MailNotifyAddress = os.Getenv("MAIL_EMAIL_ADDRESS")
	MailApiUrl = os.Getenv("MAIL_API_URL")
	MailAPIKey = os.Getenv("MAIL_API_KEY")
	MailFromAddress = os.Getenv("MAIL_FROM_ADDRESS")
}
