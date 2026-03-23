package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/0nF1REy/go-workspace/services/bubble_tea/01_health_checker/internal/model"
)

func waitResult(ch chan model.Result) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}

func (d Dashboard) Init() tea.Cmd {
	return waitResult(d.stream)
}

func (d Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case model.Result:
		d.results[msg.URL] = msg
		if _, ok := d.history[msg.URL]; !ok {
			d.history[msg.URL] = []float64{}
			if d.selected == "" {
				d.selected = msg.URL // auto-select first service
			}
		}

		d.history[msg.URL] = append(d.history[msg.URL], float64(msg.Latency.Milliseconds()))
		if len(d.history[msg.URL]) > 30 {
			d.history[msg.URL] = d.history[msg.URL][1:]
		}
		return d, waitResult(d.stream)
	case tea.KeyMsg:
		keys := make([]string, 0, len(d.results))
		for k := range d.results {
			keys = append(keys, k)
		}

		switch msg.String() {
		case "q":
			return d, tea.Quit
		case "up":
			d.selected = prevKey(keys, d.selected)
		case "down":
			d.selected = nextKey(keys, d.selected)
		}

	}
	return d, nil
}

func nextKey(keys []string, current string) string {
	for i, k := range keys {
		if k == current {
			return keys[(i+1)%len(keys)]
		}
	}
	return current
}

func prevKey(keys []string, current string) string {
	for i, k := range keys {
		if k == current {
			if i == 0 {
				return keys[len(keys)-1]
			}
			return keys[i-1]
		}
	}
	return current
}
