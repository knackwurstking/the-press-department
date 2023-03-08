package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Background for the game (just some shit with grey)
type Background struct{}

func NewBackground() *Background {
	return &Background{}
}

func (*Background) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 180, 255})
}
