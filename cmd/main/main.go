package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/moxuz/price-protection-notifier/check"
	"github.com/moxuz/price-protection-notifier/config"
	"github.com/moxuz/price-protection-notifier/db"
	"github.com/moxuz/price-protection-notifier/email"
)

func handler() error {
	config.Init()
	d, err := db.NewDB()
	if err != nil {
		return err
	}
	cr := check.NewRunner(d)
	checks, err := cr.All()
	if err != nil {
		return err
	}
	for r := range checks {
		fmt.Println(fmt.Sprintf("Checked %s price is %.2f change is %t", r.URL, r.Price, r.Changed))
		if r.Error != nil {
			fmt.Println(fmt.Sprintf("Error %s", r.Error))
			err = email.SendError(config.MailNotifyAddress, r.URL, r.Error)
		}
		if r.Changed {
			fmt.Println(fmt.Sprintf("Price changed %.2f", r.Price))
			err = email.SendPriceChange(config.MailNotifyAddress, r.URL, r.Price)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
