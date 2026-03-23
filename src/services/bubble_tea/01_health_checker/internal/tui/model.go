package tui

import "github.com/0nF1REy/go-workspace/services/bubble_tea/01_health_checker/internal/model"

type Dashboard struct {
	results  map[string]model.Result
	history  map[string][]float64
	selected string
	stream   chan model.Result
}
