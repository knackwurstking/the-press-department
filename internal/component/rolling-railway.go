// TODO: Rename to RollingRailway
package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type RollingRailway struct {
	data                      *RollingRailwayData
	rolls                     []Coord
	screenWidth, screenHeight float64
}

func NewRollingRailway(data *RollingRailwayData) Component[RollingRailwayData] {
	c := &RollingRailway{
		data:  data,
		rolls: make([]Coord, 0),
	}

	return c
}

func (c *RollingRailway) Layout(outsideWidth, outsideHeight int) (int, int) {
	c.screenWidth = float64(outsideWidth)
	c.screenHeight = float64(outsideHeight)

	return outsideWidth, outsideHeight
}

func (c *RollingRailway) Update() error {
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

func (c *RollingRailway) Draw(screen *ebiten.Image) {
	for i := 0; i < len(c.rolls); i++ {
		c.data.Sprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}
}

func (c *RollingRailway) Data() *RollingRailwayData {
	return c.data
}

type RollingRailwayData struct {
	Sprite     *RollSprite
	Scale      *float64
	Hz         float64
	HzMultiply float64

	X, Y float64

	// Update fields
	rSum, r, x, y, size float64
}

func (c *RollingRailwayData) SetUpdateData(r, x, y, size float64) {
	c.r = r
	c.x = x
	c.y = y
	c.size = size
}

func (c *RollingRailwayData) SetSprite() {
	c.rSum += c.r
	w, _ := c.Sprite.GetAssetSize()
	if c.rSum >= w {
		c.Sprite.NextSprite()
		c.rSum = 0
	}
}

func (c *RollingRailwayData) GetHeight() float64 {
	_, h := c.Sprite.GetAssetSize()
	return h
}
