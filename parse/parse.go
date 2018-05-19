package parse

import (
	"io"

	"regexp"

	"strconv"

	"errors"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrNoPrice             = errors.New("unable to find price")
	priceRegex             = `[-+]?\d*\.\d+|\d+`
	possiblePriceSelectors = []string{
		".prodprice",
		"#priceblock_ourprice"}
)

func GetPrice(body io.ReadCloser) (float64, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return 0, err
	}
	var found bool
	var price float64
	for _, s := range possiblePriceSelectors {
		selector := doc.Find(s).First()
		priceString := selector.Find("span").Text()
		if priceString != "" {
			re := regexp.MustCompile(priceRegex)
			priceList := re.FindAllString(priceString, 1)
			if len(priceList) > 0 {
				price, err = strconv.ParseFloat(priceList[0], 64)
				if err != nil {
					return 0, nil
				}
				found = true
				break
			}
		}

	}
	if !found {
		return 0, ErrNoPrice
	}
	return price, nil
}
