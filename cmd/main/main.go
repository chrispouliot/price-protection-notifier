package main

import (
	"fmt"

	"github.com/moxuz/price-protection-notifier/check"
	"github.com/moxuz/price-protection-notifier/db"
	"github.com/techdroplabs/dyspatch/web-pilot-idp/config"
)

func handler() {
	config.Init()
	d, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	cr := check.NewRunner(d)
	checks, err := cr.All()
	if err != nil {
		panic(err)
	}
	for r := range checks {
		fmt.Println(fmt.Sprintf("Checked %s price is %f change is %t", r.URL, r.Price, r.Changed))
		if r.Error != nil {
			fmt.Println(fmt.Sprintf("Error %s", r.Error))
		}
		if r.Changed {
			fmt.Println(fmt.Sprintf("Price changed %.2f", r.Price))
		}
	}
}

func main() {
	//lambda.Start(handler)
	handler()
}
