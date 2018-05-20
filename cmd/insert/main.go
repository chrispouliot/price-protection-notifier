package main

import (
	"os"

	"strconv"

	"github.com/moxuz/price-protection-notifier/config"
	"github.com/moxuz/price-protection-notifier/db"
)

func main() {
	config.Init()
	args := os.Args[1:]
	if len(args) < 2 {
		panic("Invalid arguments")
	}
	url, priceStr := args[0], args[1]
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		panic(err)
	}

	d, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	err = d.Insert(url, price)
	if err != nil {
		panic(err)
	}
}
