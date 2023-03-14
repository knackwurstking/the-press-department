package game

import (
	"bytes"
	"image"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ImageRollAsset *ebiten.Image
)

func init() {
	// Roll asset (sprit 1)
	img, _, err := image.Decode(bytes.NewReader(images.RollAsset))
	if err != nil {
		panic(err)
	}
	ImageRollAsset = ebiten.NewImageFromImage(img)
}

type RollAsset struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions

	scale *float64
}

func NewRollAsset(scale *float64, roll *ebiten.Image) Roll {
	return Roll{
		Image: roll,
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		scale: scale,
	}
}

func (r RollAsset) Draw(screen *ebiten.Image, x, y float64) {
	r.Options.GeoM.Reset()
	r.Options.GeoM.Scale(*r.scale, *r.scale)
	// TODO: draw 1-6th part from the image

	screen.DrawImage(r.Image, r.Options)
}

func (r RollAsset) GetHeight() float64 {
	_, h := r.Image.Size()
	return float64(h) * *r.scale
}

func (r RollAsset) GetWidth() float64 {
	w, _ := r.Image.Size()
	return float64(w) * *r.scale / 6
}
