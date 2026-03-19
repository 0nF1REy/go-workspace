package src

import (
	"fmt"
	"image/color"

	"github.com/0nF1REy/go-workspace/games/01_space_game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Timer
	score            int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

// Roda em 60 FPS
// Responsável por atualizar a lógica do jogo
// 60 X por segundo
// 1 x rodando
func (g *Game) Update() error {
	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			fmt.Println("Você perdeu!")
			g.Reset()
		}
	}

	for i := len(g.meteors) - 1; i >= 0; i-- {
		m := g.meteors[i]

		for j := len(g.lasers) - 1; j >= 0; j-- {
			l := g.lasers[j]

			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score++
				break
			}
		}
	}

	return nil
}

// Responsável por desenhar objetos na tela
// 60 X por segundo
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(20, 100)
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.FontUi, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
}
