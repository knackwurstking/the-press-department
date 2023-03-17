package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Coord struct {
	X, Y float64
}

type ConveyorData struct {
	Sprite     *RollSprite
	Scale      *float64
	Hz         float64
	HzMultiply float64

	X, Y float64

	// Update fields
	rSum, r, x, y, size float64
}

func (c *ConveyorData) SetUpdateData(r, x, y, size float64) {
	c.r = r
	c.x = x
	c.y = y
	c.size = size
}

func (c *ConveyorData) SetSprite() {
	c.rSum += c.r
	w, _ := c.Sprite.GetAssetSize()
	if c.rSum >= w {
		c.Sprite.NextSprite()
		c.rSum = 0
	}
}

func (c *ConveyorData) GetHeight() float64 {
	_, h := c.Sprite.GetAssetSize()
	return h
}

type Conveyor struct {
	data                      *ConveyorData
	rolls                     []Coord
	screenWidth, screenHeight float64
}

func NewConveyor(data *ConveyorData) *Conveyor {
	c := &Conveyor{
		data:  data,
		rolls: make([]Coord, 0),
	}

	return c
}

func (c *Conveyor) Layout(outsideWidth, outsideHeight int) (int, int) {
	c.screenWidth = float64(outsideWidth)
	c.screenHeight = float64(outsideHeight)

	return outsideWidth, outsideHeight
}

func (c *Conveyor) Update() error {
	c.data.X = c.data.x
	c.data.Y = c.data.y

	c.data.SetSprite()
	w, _ := c.data.Sprite.GetAssetSize()
	padding := w * 3

	c.rolls = make([]Coord, 0)
	for p := c.data.X; p <= c.data.size; p += (w + padding) {
		c.rolls = append(c.rolls, Coord{X: float64(p), Y: c.data.Y})
	}

	return nil
}

func (c *Conveyor) Draw(screen *ebiten.Image) {
	for i := 0; i < len(c.rolls); i++ {
		c.data.Sprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}
}

func (c *Conveyor) SetData(data *ConveyorData) {
	c.data = data
}

func (c *Conveyor) GetData() *ConveyorData {
	return c.data
}
