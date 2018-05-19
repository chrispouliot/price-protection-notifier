package parse

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func getHtmlReader(path string) (io.ReadCloser, error) {
	b, err := os.Open(path)
	return b, err
}

func TestParseBestBuy(t *testing.T) {
	r, err := getHtmlReader("fixtures/bestbuy.html")
	require.NoError(t, err)

	price, err := GetPrice(r)
	require.NoError(t, err)
	require.Equal(t, 869.99, price)
}

func TestParseAmazon(t *testing.T) {
	r, err := getHtmlReader("fixtures/amazon.html")
	require.NoError(t, err)

	price, err := GetPrice(r)
	require.NoError(t, err)
	require.Equal(t, 249.95, price)
}

func TestParseWayfair(t *testing.T) {
	r, err := getHtmlReader("fixtures/wayfair.html")
	require.NoError(t, err)

	price, err := GetPrice(r)
	require.NoError(t, err)
	require.Equal(t, 310.99, price)
}

func TestParseNoPrice(t *testing.T) {
	r, err := getHtmlReader("fixtures/noprice.html")
	require.NoError(t, err)

	_, err = GetPrice(r)
	require.Error(t, err)
	require.Equal(t, err, ErrNotFound)
}
