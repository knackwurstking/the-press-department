package game

import "github.com/hajimehoshi/ebiten/v2"

type Conveyor struct {
	Scale *float64
}

func NewConveyor(scale *float64) *Conveyor {
	// TODO: need some game assets first (to simulate a running rb)
	return &Conveyor{
		Scale: scale,
	}
}

func (c *Conveyor) Draw(screen *ebiten.Image, x, y, width, height float64) {
}

func (c *Conveyor) GetHeight() float64 {
	return 0
}
