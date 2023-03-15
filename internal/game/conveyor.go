package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Conveyor struct {
	hz         float64
	hzMultiply float64
	rolls      []Coord
	scale      *float64
	sprite     *Roll
	r          float64
}

func NewConveyor(scale *float64, hzMultiply float64) Conveyor {
	return Conveyor{
		hzMultiply: hzMultiply,
		rolls:      make([]Coord, 0),
		scale:      scale,
		sprite:     NewRoll(scale),
	}
}

func (c *Conveyor) Draw(screen *ebiten.Image) {
	for i := 0; i < len(c.rolls); i++ {
		c.sprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}
}

func (c *Conveyor) Update(r, x, y, size float64) {
	c.SetSprite(r)
	w, _ := c.sprite.GetAssetSize()
	padding := w * 3

	c.rolls = make([]Coord, 0)
	for p := x; p <= size; p += (w + padding) {
		c.rolls = append(c.rolls, Coord{X: float64(p), Y: y})
	}
}

func (c *Conveyor) GetHeight() float64 {
	_, h := c.sprite.GetAssetSize()
	return h
}

func (c *Conveyor) SetSprite(r float64) {
	c.r += r
	w, _ := c.sprite.GetAssetSize()
	if c.r >= w {
		c.sprite.NextSprite()
		c.r = 0
	}
}
