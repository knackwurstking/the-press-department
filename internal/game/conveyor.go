package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Conveyor struct {
	hz               float64
	hzMultiply       float64
	rollTypes        [3]Sprite
	rolls            []SpriteCoord
	scale            *float64
	sprite           Sprite
	lastSpriteRender time.Time
	nextSprite       int
}

func NewConveyor(scale *float64, hzMultiply float64) Conveyor {
	return Conveyor{
		hzMultiply: hzMultiply,
		rollTypes: [3]Sprite{
			NewRoll(scale, ImageRoll2),
			NewRoll(scale, ImageRoll1),
			NewRoll(scale, ImageRoll0),
		},
		rolls: make([]SpriteCoord, 0),
		scale: scale,
	}
}

func (c *Conveyor) Draw(screen *ebiten.Image) {
	for i := 0; i < len(c.rolls); i++ {
		c.sprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}
}

func (c *Conveyor) Update(prev, current time.Time, x, y, size float64) {
	c.SetSprite(prev, current)
	w := c.sprite.GetWidth()
	padding := w * 3

	c.rolls = make([]SpriteCoord, 0)
	for p := x; p <= size; p += (w + padding) {
		c.rolls = append(c.rolls, SpriteCoord{X: float64(p), Y: y})
	}
}

func (c *Conveyor) GetHeight() float64 {
	return c.rollTypes[0].GetHeight()
}

func (c *Conveyor) SetSprite(prev, current time.Time) {
	if current.Sub(c.lastSpriteRender).Seconds()*(c.hz*c.hzMultiply)*(*c.scale*10) > (60 / (c.hz)) {
		c.lastSpriteRender = current

		c.nextSprite += 1
		if c.nextSprite >= len(c.rollTypes) {
			c.nextSprite = 0
		}
	}

	c.sprite = c.rollTypes[c.nextSprite]
}
