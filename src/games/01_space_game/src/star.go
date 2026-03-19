package src

import (
	"math/rand"

	"github.com/0nF1REy/go-workspace/games/01_space_game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Star struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewStar() *Star {
	// Escolhe uma imagem de estrela aleatória dos assets
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]

	// Velocidade variada para dar um efeito de profundidade (parallax)
	speed := (rand.Float64() * 2) + 1 // Entre 1 e 3

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -20, // Começa um pouco acima da tela
	}

	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (s *Star) Update() {
	s.position.Y += s.speed
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.position.X, s.position.Y)

	// Diminui um pouco a opacidade para as estrelas ficarem mais ao fundo
	op.ColorScale.Scale(1, 1, 1, 0.7)

	screen.DrawImage(s.image, op)
}

// Verifica se a estrela saiu da tela para podermos removê-la da memória
func (s *Star) IsOffScreen() bool {
	return s.position.Y > screenHeight
}
