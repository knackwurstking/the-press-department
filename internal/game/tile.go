package game

import (
	_ "image/jpeg"

	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"the-press-department/internal/images"
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
	_, h := t.Image.Size()
	return float64(h) * *t.scale
}

func (t *Tile) GetWidth() float64 {
	w, _ := t.Image.Size()
	return float64(w) * *t.scale
}
