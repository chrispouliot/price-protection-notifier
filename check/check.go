package check

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/moxuz/price-protection-notifier/db"
	"github.com/moxuz/price-protection-notifier/parse"
)

type Runner struct {
	d *db.DB
}

type Result struct {
	Error   error
	Changed bool
	Price   float64
	URL     string
}

func NewRunner(db *db.DB) *Runner {
	return &Runner{d: db}
}

func (r *Runner) All() (<-chan *Result, error) {
	var wg sync.WaitGroup
	rChan := make(chan *Result)
	checks, err := r.d.GetAll()
	if err != nil {
		return nil, err
	}
	for _, c := range checks {
		wg.Add(1)
		go r.run(rChan, c, &wg)
	}
	go func() {
		wg.Wait()
		close(rChan)
	}()
	return rChan, nil
}

func (r *Runner) run(c chan *Result, check *db.Check, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(check.URL)
	if err != nil {
		c <- &Result{Error: err}
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		c <- &Result{Error: errors.New(fmt.Sprintf(
			"Request returned status of %d: %s", resp.StatusCode, resp.Status,
		))}
		return
	}

	price, err := parse.GetPrice(resp.Body)
	if err != nil {
		c <- &Result{Error: err}
		return
	}
	c <- &Result{URL: check.URL, Changed: price < check.LastPrice, Price: price}
}
