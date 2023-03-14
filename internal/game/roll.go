package game

import (
	"bytes"
	"image"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ImageRoll0 *ebiten.Image
	ImageRoll1 *ebiten.Image
	ImageRoll2 *ebiten.Image
)

func init() {
	// Roll asset (sprit 1)
	img, _, err := image.Decode(bytes.NewReader(images.Roll0))
	if err != nil {
		panic(err)
	}
	ImageRoll0 = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 2)
	img, _, err = image.Decode(bytes.NewReader(images.Roll1))
	if err != nil {
		panic(err)
	}
	ImageRoll1 = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 3)
	img, _, err = image.Decode(bytes.NewReader(images.Roll2))
	if err != nil {
		panic(err)
	}
	ImageRoll2 = ebiten.NewImageFromImage(img)
}

type Roll struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions

	scale *float64
}

func NewRoll(scale *float64, roll *ebiten.Image) Roll {
	return Roll{
		Image: roll,
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		scale: scale,
	}
}

func (r *Roll) Draw(screen *ebiten.Image, x, y float64) {
	r.Options.GeoM.Reset()
	r.Options.GeoM.Scale(*r.scale, *r.scale)
	r.Options.GeoM.Translate(x, y)

	screen.DrawImage(r.Image, r.Options)
}

func (r *Roll) GetHeight() float64 {
	_, h := r.Image.Size()
	return float64(h) * *r.scale
}

func (r *Roll) GetWidth() float64 {
	w, _ := r.Image.Size()
	return float64(w) * *r.scale
}
