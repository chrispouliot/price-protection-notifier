package main

import (
	"os"

	"fmt"

	"github.com/moxuz/price-protection-notifier/config"
	"github.com/moxuz/price-protection-notifier/db"
)

func main() {
	config.Init()
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Invalid arguments")
	}

	url := args[0]

	d, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	numRemoved, err := d.Delete(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Removed %d items", numRemoved))
}
