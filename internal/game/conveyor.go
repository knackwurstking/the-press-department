package game

import "github.com/hajimehoshi/ebiten/v2"

type Conveyor struct {
	rollTypes [3]Roll
	rolls     []Roll
	scale     *float64
}

func NewConveyor(scale *float64) Conveyor {
	return Conveyor{
		rollTypes: [3]Roll{
			NewRoll(scale, ImageRoll0),
			NewRoll(scale, ImageRoll1),
			NewRoll(scale, ImageRoll2),
		},
		rolls: make([]Roll, 0),
		scale: scale,
	}
}

func (c *Conveyor) Draw(screen *ebiten.Image, x, y, width, height float64) {
	// TODO: draw the rollers along the x axis
}

func (c *Conveyor) GetHeight() float64 {
	return c.rollTypes[0].GetHeight()
}
