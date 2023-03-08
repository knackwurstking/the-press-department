package game

import (
	_ "image/jpeg"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	UrbanDoveActive = NewTileAsset("assets/tiles/urban-active-urban-dove-active.jpg")
)

type TileAsset struct {
	Path  string
	Image *ebiten.Image
}

func NewTileAsset(path string) *TileAsset {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}

	return &TileAsset{
		Path:  path,
		Image: img,
	}
}

type Tile struct {
	Src    string
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
