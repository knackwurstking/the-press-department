package game

import "github.com/hajimehoshi/ebiten/v2"

// NOTE:
// 	- sprite used for `Roll` (Conveyor)
// 	- maybe for tiles? (Engines)

type Sprite interface {
	Draw(screen *ebiten.Image, x, y float64)
	GetHeight() float64
	GetWidth() float64
}

type SpriteCoord struct {
	Sprite Sprite
	X, Y   float64
}
