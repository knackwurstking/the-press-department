package game

import (
	"bytes"
	"image"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ImageRoll [8]*ebiten.Image
)

func init() {
	// Roll asset (sprit 0)
	img, _, err := image.Decode(bytes.NewReader(images.Roll0))
	if err != nil {
		panic(err)
	}
	ImageRoll[0] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 1)
	img, _, err = image.Decode(bytes.NewReader(images.Roll1))
	if err != nil {
		panic(err)
	}
	ImageRoll[1] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 2)
	img, _, err = image.Decode(bytes.NewReader(images.Roll2))
	if err != nil {
		panic(err)
	}
	ImageRoll[2] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 3)
	img, _, err = image.Decode(bytes.NewReader(images.Roll3))
	if err != nil {
		panic(err)
	}
	ImageRoll[3] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 4)
	img, _, err = image.Decode(bytes.NewReader(images.Roll4))
	if err != nil {
		panic(err)
	}
	ImageRoll[4] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 5)
	img, _, err = image.Decode(bytes.NewReader(images.Roll5))
	if err != nil {
		panic(err)
	}
	ImageRoll[5] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 6)
	img, _, err = image.Decode(bytes.NewReader(images.Roll6))
	if err != nil {
		panic(err)
	}
	ImageRoll[6] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 7)
	img, _, err = image.Decode(bytes.NewReader(images.Roll7))
	if err != nil {
		panic(err)
	}
	ImageRoll[7] = ebiten.NewImageFromImage(img)
}

type Roll struct {
	imageIndex int
	Options    *ebiten.DrawImageOptions

	scale *float64
}

func NewRoll(scale *float64) *Roll {
	return &Roll{
		imageIndex: 0,
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

	screen.DrawImage(ImageRoll[r.imageIndex], r.Options)
}

func (r *Roll) GetAssetSize() (width float64, height float64) {
	w, h := ImageRoll[0].Size()
	return float64(w) * *r.scale, float64(h) * *r.scale
}

func (r *Roll) NextSprite() {
	r.imageIndex += 1

	if r.imageIndex >= len(ImageRoll) {
		r.imageIndex = 0
	}
}
