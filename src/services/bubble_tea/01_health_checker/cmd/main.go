package main

import (
	"github.com/0nF1REy/go-workspace/services/bubble_tea/01_health_checker/internal/checker"
	"github.com/0nF1REy/go-workspace/services/bubble_tea/01_health_checker/internal/tui"
)

func main() {
	stream := checker.Start()
	app := tui.New(stream)
	app.Run()
}
