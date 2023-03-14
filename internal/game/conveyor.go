package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Conveyor struct {
	hz         float64
	hzMultiply float64
	rollTypes  [3]Sprite
	rolls      []SpriteCoord
	scale      *float64
}

func NewConveyor(scale *float64, hzMultiply float64) Conveyor {
	return Conveyor{
		hzMultiply: hzMultiply,
		rollTypes: [3]Sprite{
			NewRoll(scale, ImageRoll0),
			NewRoll(scale, ImageRoll1),
			NewRoll(scale, ImageRoll2),
		},
		rolls: make([]SpriteCoord, 0),
		scale: scale,
	}
}

func (c *Conveyor) Draw(screen *ebiten.Image) {
	for i := 0; i < len(c.rolls); i++ {
		c.rolls[i].Sprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}
}

func (c *Conveyor) Update(prev, current time.Time, x, y, size float64) {
	sprite := c.getSprite(prev, current)
	w := sprite.GetWidth()
	padding := w * 3

	c.rolls = make([]SpriteCoord, 0)
	for position := x; x <= size; position += w + padding {
		c.rolls = append(c.rolls, SpriteCoord{Sprite: sprite, X: position, Y: y})
	}
}

func (c *Conveyor) GetHeight() float64 {
	return c.rollTypes[0].GetHeight()
}

func (c *Conveyor) getSprite(prev, current time.Time) Sprite {
	// TODO: get sprite sheet based on the Engines `Hz` and the prev and current values
	return c.rollTypes[0]
}
