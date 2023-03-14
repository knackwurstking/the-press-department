package game

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ImageTile          *ebiten.Image
	ImageTileWithCrack *ebiten.Image
)

func init() {
	// Tile
	img, _, err := image.Decode(bytes.NewReader(images.Tile))
	if err != nil {
		panic(err)
	}
	ImageTile = ebiten.NewImageFromImage(img)

	// TileWithCrack
	img, _, err = image.Decode(bytes.NewReader(images.TileWithCrack))
	if err != nil {
		panic(err)
	}
	ImageTileWithCrack = ebiten.NewImageFromImage(img)
}

type Tile struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
	X       float64

	scale *float64
}

func NewTile(scale *float64, tile *ebiten.Image) Tile {
	return Tile{
		Image: tile,
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		X:     0,
		scale: scale,
	}
}

func (t *Tile) Draw(screen *ebiten.Image, x, y float64) {
	t.Options.GeoM.Reset()
	t.Options.GeoM.Scale(*t.scale, *t.scale)
	t.Options.GeoM.Translate(x, y)

	screen.DrawImage(t.Image, t.Options)
}

func (t *Tile) GetHeight() float64 {
	_, height := t.Image.Size()
	return float64(height) * *t.scale
}

func (t *Tile) GetWidth() float64 {
	width, _ := t.Image.Size()
	return float64(width) * *t.scale
}
