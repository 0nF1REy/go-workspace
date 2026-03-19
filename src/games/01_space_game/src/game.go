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
	stars            []*Star
	meteorSpawnTimer *Timer
	starSpawnTimer   *Timer
	score            int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
		starSpawnTimer:   NewTimer(10), // Estrelas aparecem com mais frequência
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

func (g *Game) Update() error {
	// 1. Atualizar Estrelas
	g.starSpawnTimer.Update()
	if g.starSpawnTimer.IsReady() {
		g.starSpawnTimer.Reset()
		g.stars = append(g.stars, NewStar())
	}

	for i := len(g.stars) - 1; i >= 0; i-- {
		g.stars[i].Update()
		if g.stars[i].IsOffScreen() {
			g.stars = append(g.stars[:i], g.stars[i+1:]...)
		}
	}

	// 2. Atualizar Jogador e Lasers
	g.player.Update()
	for _, l := range g.lasers {
		l.Update()
	}

	// 3. Atualizar Meteoros
	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	// Colisões
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

func (g *Game) Draw(screen *ebiten.Image) {
	// Desenhando as estrelas PRIMEIRO para ficarem atrás de tudo
	for _, s := range g.stars {
		s.Draw(screen)
	}

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
	g.stars = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
	g.starSpawnTimer.Reset()
	g.score = 0
}
