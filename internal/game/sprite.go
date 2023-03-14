package game

import "github.com/hajimehoshi/ebiten/v2"

type Sprite interface {
	Draw(screen *ebiten.Image, x, y float64)
	GetHeight() float64
	GetWidth() float64
}

type SpriteCoord struct {
	Sprite Sprite
	X, Y   float64
}
