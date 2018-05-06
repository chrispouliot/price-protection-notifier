package check

import (
	"github.com/moxuz/price-protection-notifier/db"
	"sync"
)

type CheckRunner struct {
	d *db.DB
}

type Result struct {
	Changed bool
	URL     string
}

func NewRunner(db *db.DB) *CheckRunner {
	return &CheckRunner{d: db}
}

func (r *CheckRunner) RunAll() <-chan *Result {
	var wg sync.WaitGroup
	rChan := make(chan *Result)
	checks, err := r.d.GetAll()
	if err != nil {
		// do something with err
	}
	for _, c := range checks {
		wg.Add(1)
		go run(rChan, c, wg)
	}
	go func() {
		wg.Wait()
		close(rChan)
	}()
	return rChan
}

func run(c chan *Result, check *db.Check, wg sync.WaitGroup) {
	// do check
	c <- &Result{URL: "asd", Changed: false}
	wg.Done()
}
