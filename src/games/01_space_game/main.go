package main

import (
	"github.com/0nF1REy/go-workspace/games/01_space_game/src"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := src.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
