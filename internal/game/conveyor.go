package game

import "github.com/hajimehoshi/ebiten/v2"

type Conveyor struct {
	scale *float64
}

func NewConveyor(scale *float64) Conveyor {
	// TODO: need some game assets first (to simulate a running rb)
	return Conveyor{
		scale: scale,
	}
}

func (c *Conveyor) Draw(screen *ebiten.Image, x, y, width, height float64) {
	// TODO: draw the rollers along the x axis
}

func (c *Conveyor) GetHeight() float64 {
	// TODO: get height from image assets (only rollers for now)
	return 0
}
