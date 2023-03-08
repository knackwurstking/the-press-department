package game

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"the-press-department/internal/images"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	urbanDoveActive *ebiten.Image
)

func init() {
	var err error
	var img image.Image
	img, _, err = image.Decode(bytes.NewReader(images.UrbanDoveActive))
	if err != nil {
		panic(err)
	}

	urbanDoveActive = ebiten.NewImageFromImage(img)
}

type Tile struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
	Color   color.Color
	Scale   *float64
	X       float64
}

func NewTile(scale *float64) *Tile {
	return &Tile{
		Image: ebiten.NewImageFromImage(urbanDoveActive),
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		Color: color.RGBA{0, 0, 0, 255},
		Scale: scale,
		X:     0,
	}
}

func (t *Tile) Draw(screen *ebiten.Image, x, y float64) {
	t.Options.GeoM.Reset()
	t.Options.GeoM.Scale(*t.Scale, *t.Scale)
	t.Options.GeoM.Translate(x, y)

	screen.DrawImage(t.Image, t.Options)
}

func (t *Tile) GetHeight() float64 {
	_, height := t.Image.Size()
	return float64(height) * *t.Scale
}

func (t *Tile) GetWidth() float64 {
	width, _ := t.Image.Size()
	return float64(width) * *t.Scale
}
