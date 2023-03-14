package game

import "github.com/hajimehoshi/ebiten/v2"

type Coord struct {
	x, y float64
}

type Conveyor struct {
	rollTypes [3]Roll
	rolls     []Coord
	scale     *float64

	x, y, width, height float64
}

func NewConveyor(scale *float64, x, y, width, height float64) Conveyor {
	return Conveyor{
		rollTypes: [3]Roll{
			NewRoll(scale, ImageRoll0),
			NewRoll(scale, ImageRoll1),
			NewRoll(scale, ImageRoll2),
		},
		rolls:  make([]Coord, 0),
		scale:  scale,
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func (c *Conveyor) Draw(screen *ebiten.Image) {
	sprite := c.getSprite()
	for i := 0; i < len(c.rolls); i++ {
		sprite.Draw(screen, c.rolls[i].x, c.rolls[i].y)
	}
}

func (c *Conveyor) Update() {
	// TODO: fill c.rolls with rolls
}

func (c *Conveyor) GetHeight() float64 {
	return c.rollTypes[0].GetHeight()
}

func (c *Conveyor) getSprite() Roll {
	return c.rollTypes[0]
}
