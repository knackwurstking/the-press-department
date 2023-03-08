package game

import "image/color"

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
