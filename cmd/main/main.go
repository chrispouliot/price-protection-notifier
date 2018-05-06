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
	for r := range cr.RunAll() {
		if r.Changed {
			// Notify!
		}
	}
}

func main() {
	lambda.Start(handler)
}
