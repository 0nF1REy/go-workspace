package model

import "time"

type Result struct {
	URL     string
	Status  string
	Latency time.Duration
	Error   string
	Checked time.Time
}
