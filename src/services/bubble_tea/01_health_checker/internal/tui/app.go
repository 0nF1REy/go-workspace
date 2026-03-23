package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/0nF1REy/go-workspace/services/bubble_tea/01_health_checker/internal/model"
)

type App struct {
	stream chan model.Result
}

func New(stream chan model.Result) *App {
	return &App{stream: stream}
}

func (a *App) Run() {
	m := Dashboard{
		results: make(map[string]model.Result),
		history: make(map[string][]float64),
		stream:  a.stream,
	}
	p := tea.NewProgram(m)
	p.Run()
}
