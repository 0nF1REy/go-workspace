package checker

import (
	"context"
	"net/http"
	"time"

	"github.com/0nF1REy/go-workspace/services/bubble_tea/01_health_checker/internal/model"
)

func checkUrl(ctx context.Context, url string) model.Result {
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return model.Result{
			URL:     url,
			Status:  "DOWN",
			Latency: 0,
			Error:   err.Error(),
			Checked: time.Now(),
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	latency := time.Since(start)

	if err != nil {
		return model.Result{
			URL:     url,
			Status:  "DOWN",
			Latency: latency,
			Error:   err.Error(),
			Checked: time.Now(),
		}
	}
	defer resp.Body.Close()

	status := "UP"
	if resp.StatusCode >= 400 {
		status = "DOWNGRADE"
	}

	return model.Result{
		URL:     url,
		Status:  status,
		Latency: latency,
		Checked: time.Now(),
	}
}
