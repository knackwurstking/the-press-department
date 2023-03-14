package game

import "github.com/hajimehoshi/ebiten/v2"

type Conveyor struct {
	rollTypes [3]Sprite
	rolls     []SpriteCoord
	scale     *float64

	x, y, width, height float64
}

func NewConveyor(scale *float64, x, y, width, height float64) Conveyor {
	return Conveyor{
		rollTypes: [3]Sprite{
			NewRoll(scale, ImageRoll0),
			NewRoll(scale, ImageRoll1),
			NewRoll(scale, ImageRoll2),
		},
		rolls:  make([]SpriteCoord, 0),
		scale:  scale,
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func (c *Conveyor) Draw(screen *ebiten.Image) {
	for i := 0; i < len(c.rolls); i++ {
		c.rolls[i].Sprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}
}

func (c *Conveyor) Update() {
	sprite := c.getSprite()
	w := sprite.GetWidth()
	p := w * 3 // padding

	c.rolls = make([]SpriteCoord, 0)
	for x := c.x; x <= c.width; x += w + p {
		c.rolls = append(c.rolls, SpriteCoord{Sprite: sprite, X: x, Y: c.y})
	}
}

func (c *Conveyor) GetHeight() float64 {
	return c.rollTypes[0].GetHeight()
}

func (c *Conveyor) getSprite() Sprite {
	return c.rollTypes[0]
}
