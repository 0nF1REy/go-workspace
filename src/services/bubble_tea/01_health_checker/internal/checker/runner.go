package checker

import (
	"context"
	"sync"
	"time"

	"github.com/0nF1REy/go-workspace/services/bubble_tea/01_health_checker/internal/model"
)

func runChecks(urls []string, out chan<- model.Result) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			out <- checkUrl(ctx, u)
		}(url)
	}
	wg.Wait()
}
