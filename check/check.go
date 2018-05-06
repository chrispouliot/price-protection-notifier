package check

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/moxuz/price-protection-notifier/db"
)

type CheckRunner struct {
	d *db.DB
}

type Result struct {
	Error   error
	Changed bool
	URL     string
}

func NewRunner(db *db.DB) *CheckRunner {
	return &CheckRunner{d: db}
}

func (r *CheckRunner) RunAll() (<-chan *Result, error) {
	var wg sync.WaitGroup
	rChan := make(chan *Result)
	checks, err := r.d.GetAll()
	if err != nil {
		return nil, err
	}
	for _, c := range checks {
		wg.Add(1)
		go r.run(rChan, c, wg)
	}
	go func() {
		wg.Wait()
		close(rChan)
	}()
	return rChan, nil
}

func (r *CheckRunner) run(c chan *Result, check *db.Check, wg sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(check.URL)
	if err != nil {
		c <- &Result{Error: err}
		return
	}
	if resp.StatusCode != 200 {
		c <- &Result{Error: errors.New(fmt.Sprintf(
			"Request returned status of %d: %s", resp.StatusCode, resp.Status,
		))}
		return
	}
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- &Result{Error: err}
		return
	}
	defer resp.Body.Close()
	price := PriceFromHTML(html)
	c <- &Result{URL: check.URL, Changed: price < check.LastPrice}
}
