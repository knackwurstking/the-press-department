package component

import (
	"bytes"
	"image"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ImageRollAsset [8]*ebiten.Image
)

func init() {
	// Roll asset (sprit 0)
	img, _, err := image.Decode(bytes.NewReader(images.Roll0))
	if err != nil {
		panic(err)
	}
	ImageRollAsset[0] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 1)
	img, _, err = image.Decode(bytes.NewReader(images.Roll1))
	if err != nil {
		panic(err)
	}
	ImageRollAsset[1] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 2)
	img, _, err = image.Decode(bytes.NewReader(images.Roll2))
	if err != nil {
		panic(err)
	}
	ImageRollAsset[2] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 3)
	img, _, err = image.Decode(bytes.NewReader(images.Roll3))
	if err != nil {
		panic(err)
	}
	ImageRollAsset[3] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 4)
	img, _, err = image.Decode(bytes.NewReader(images.Roll4))
	if err != nil {
		panic(err)
	}
	ImageRollAsset[4] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 5)
	img, _, err = image.Decode(bytes.NewReader(images.Roll5))
	if err != nil {
		panic(err)
	}
	ImageRollAsset[5] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 6)
	img, _, err = image.Decode(bytes.NewReader(images.Roll6))
	if err != nil {
		panic(err)
	}
	ImageRollAsset[6] = ebiten.NewImageFromImage(img)

	// Roll asset (sprit 7)
	img, _, err = image.Decode(bytes.NewReader(images.Roll7))
	if err != nil {
		panic(err)
	}
	ImageRollAsset[7] = ebiten.NewImageFromImage(img)
}

type RollSprite struct {
	imageIndex int
	Options    *ebiten.DrawImageOptions

	scale *float64
}

func NewRollSprite(scale *float64) *RollSprite {
	return &RollSprite{
		imageIndex: 0,
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		scale: scale,
	}
}

func (r *RollSprite) Draw(screen *ebiten.Image, x, y float64) {
	r.Options.GeoM.Reset()
	r.Options.GeoM.Scale(*r.scale, *r.scale)
	r.Options.GeoM.Translate(x, y)

	screen.DrawImage(ImageRollAsset[r.imageIndex], r.Options)
}

func (r *RollSprite) GetAssetSize() (width float64, height float64) {
	w := ImageRollAsset[0].Bounds().Dx()
	h := ImageRollAsset[0].Bounds().Dy()
	return float64(w) * *r.scale, float64(h) * *r.scale
}

func (r *RollSprite) NextSprite() {
	r.imageIndex += 1

	if r.imageIndex >= len(ImageRollAsset) {
		r.imageIndex = 0
	}
}
