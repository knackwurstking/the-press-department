package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Background for the game (just some shit with grey)
type Background struct {
	scale float64
}

func NewBackground() *Background {
	return &Background{
		scale: DefaultScale,
	}
}

func (*Background) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 180, 255})
}

func (b *Background) GetScale() float64 {
	return b.scale
}

func (b *Background) SetScale(f float64) {
	b.scale = f
}
