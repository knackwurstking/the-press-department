package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Conveyor struct {
	hz               float64
	hzMultiply       float64
	rolls            []Coord
	scale            *float64
	sprite           *Roll
	lastSpriteRender time.Time
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

func (c *Conveyor) Update(current time.Time, x, y, size float64) {
	c.SetSprite(current)
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

func (c *Conveyor) SetSprite(current time.Time) {
	if current.Sub(c.lastSpriteRender).Seconds()*(c.hz*c.hzMultiply)*(*c.scale*10) > (60 / (c.hz)) {
		c.lastSpriteRender = current
		c.sprite.NextSprite()
	}
}
