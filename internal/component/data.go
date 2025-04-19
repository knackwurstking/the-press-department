package component

import (
	"the-press-department/internal/sprites"
	"the-press-department/internal/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type BackgroundData struct {
	Image *ebiten.Image
}

type RollerConveyorData struct {
	RollSprite *sprites.Roll

	// TODO: Ok, i hate this: X, x, Y, y
	X, Y float64

	// Update fields
	rSum, r, x, y, size float64

	tiles []tiles.Tiles
}

// TODO: Rename this method, what will be updated with this data
func (c *RollerConveyorData) SetUpdateData(r, x, y, size float64) {
	c.r = r
	c.x = x
	c.y = y
	c.size = size
}

func (c *RollerConveyorData) SetSprite() {
	c.rSum += c.r
	w, _ := c.RollSprite.GetAssetSize()
	if c.rSum >= w {
		c.RollSprite.NextSprite()
		c.rSum = 0
	}
}

func (c *RollerConveyorData) Height() float64 {
	_, h := c.RollSprite.GetAssetSize()
	return h
}

func (c *RollerConveyorData) Tiles() []tiles.Tiles {
	return c.tiles
}

type RollerConveyorUserInputData struct {
	ThrowAwayPaddingTop    float64
	ThrowAwayPaddingBottom float64
	Tiles                  []tiles.Tiles
}
