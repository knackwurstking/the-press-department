package game

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"the-press-department/internal/images"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var ()

type Tile struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
	Color   color.Color
	X       float64

	scale float64
}

func NewTile() *Tile {
	image, _, err := image.Decode(bytes.NewReader(images.UrbanDoveActive))
	if err != nil {
		panic(err)
	}

	return &Tile{
		Image: ebiten.NewImageFromImage(image),
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		Color: color.RGBA{0, 0, 0, 255},
		X:     0,
		scale: 0.2,
	}
}

func (t *Tile) Draw(screen *ebiten.Image, x, y float64) {
	t.Options.GeoM.Reset()
	t.Options.GeoM.Scale(t.scale, t.scale)
	t.Options.GeoM.Translate(x, y)

	screen.DrawImage(t.Image, t.Options)
}

func (t *Tile) GetHeight() float64 {
	_, height := t.Image.Size()
	return float64(height) * t.scale
}

func (t *Tile) GetWidth() float64 {
	width, _ := t.Image.Size()
	return float64(width) * t.scale
}

func (t *Tile) GetScale() float64 {
	return t.scale
}
