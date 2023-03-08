package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	Crack  bool // Crack holds whenever this tile has a crack or not
	Color  color.Color
	Width  float64
	Height float64
	X      float64
}

func NewTile(width, height float64) *Tile {
	return &Tile{
		Crack:  false,
		Color:  color.RGBA{0, 0, 0, 255},
		Width:  width,
		Height: height,
		X:      0,
	}
}

// TODO: adding tile assets

func (t *Tile) Draw(screen *ebiten.Image, x, y float64) {
	ebitenutil.DrawRect(
		screen,   // dst
		x,        // x - start right
		y,        // y - center
		t.Width,  // width
		t.Height, // height
		t.Color,  // color
	)
}
