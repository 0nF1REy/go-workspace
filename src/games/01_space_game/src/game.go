package src

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
	lasers []*Laser
}

func NewGame() *Game {
	g := &Game{}
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

	for _, l := range(g.lasers) {
		l.Update()
	}

	return nil
}

// Responsável por desenhar objetos na tela
// 60 X por segundo
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

		for _, l := range(g.lasers) {
		l.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}
