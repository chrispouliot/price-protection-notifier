package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/moxuz/price-protection-notifier/check"
	"github.com/moxuz/price-protection-notifier/db"
)

func handler() {
	d, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	cr := check.NewRunner(d)
	checks, err := cr.RunAll()
	if err != nil {
		panic(err)
	}
	for r := range checks {
		if r.Error != nil {
			// Notify of error!
		}
		if r.Changed {
			// Notify of price change!
		}
	}
}

func main() {
	lambda.Start(handler)
}
