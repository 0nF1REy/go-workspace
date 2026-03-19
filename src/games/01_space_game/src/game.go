package src

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
}

func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		player: player,
	}
}

// Roda em 60 FPS
// Responsável por atualizar a lógica do jogo
// 60 X por segundo
func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// Responsável por desenhar objetos na tela
// 60 X por segundo
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
