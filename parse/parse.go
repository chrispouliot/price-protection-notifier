package parse

import (
	"errors"
	"io"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrNotFound            = errors.New("unable to find price")
	priceRegex             = `[-+]?\d*\.\d+|\d+` // https://regex101.com/r/mWnOj3/2
	possiblePriceSelectors = []string{
		".prodprice",
		".ProductPricing",
		"#price",
	}
)

func GetPrice(body io.ReadCloser) (float64, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return 0, err
	}
	var found bool
	var price float64
	for _, se := range possiblePriceSelectors {
		selector := doc.Find(se).First()
		posPrice, err := searchPriceInChildren(selector.Children())
		if err == nil {
			price, found = posPrice, true
			break
		}
		if err == ErrNotFound {
			continue
		}
		return 0, err

	}
	if !found {
		return 0, ErrNotFound
	}
	return price, nil
}

func searchPriceInChildren(selection *goquery.Selection) (float64, error) {
	texts := selection.Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	for _, t := range texts {
		if t != "" {
			re := regexp.MustCompile(priceRegex)
			priceList := re.FindAllString(t, 1)
			if len(priceList) > 0 {
				price, err := strconv.ParseFloat(priceList[0], 64)
				if err != nil {
					return 0, err
				}
				return price, nil
			}
		}
	}
	return 0, ErrNotFound
}
