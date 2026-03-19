package src

import (
	"github.com/0nF1REy/go-workspace/games/01_space_game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
}

func NewPlayer() *Player {
	image := assets.PlayerSprite
	return &Player{
		image: image,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem será desenhada na tela
	op.GeoM.Translate(100, 100)

	// Desenha imagem na tela
	screen.DrawImage(p.image, op)
}
