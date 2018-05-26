package main

import (
	"fmt"

	"github.com/moxuz/price-protection-notifier/config"
	"github.com/moxuz/price-protection-notifier/db"
)

func main() {
	config.Init()
	d, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	checks, err := d.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Found %d checks", len(checks)))
	for _, check := range checks {
		fmt.Println(fmt.Sprintf("Price: %.2f URL: '%s'", check.LastPrice, check.URL))
	}
}
