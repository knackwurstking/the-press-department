package sprites

import (
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var ImageRollSprites [8]*ebiten.Image

func init() {
	ImageRollSprites[0] = images.Roll0
	ImageRollSprites[1] = images.Roll1
	ImageRollSprites[2] = images.Roll2
	ImageRollSprites[3] = images.Roll3
	ImageRollSprites[4] = images.Roll4
	ImageRollSprites[5] = images.Roll5
	ImageRollSprites[6] = images.Roll6
	ImageRollSprites[7] = images.Roll7
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

	screen.DrawImage(ImageRollSprites[r.imageIndex], r.Options)
}

func (r *Roll) GetAssetSize() (width float64, height float64) {
	w := ImageRollSprites[0].Bounds().Dx()
	h := ImageRollSprites[0].Bounds().Dy()
	return float64(w) * *r.scale, float64(h) * *r.scale
}

func (r *Roll) NextSprite() {
	r.imageIndex += 1

	if r.imageIndex >= len(ImageRollSprites) {
		r.imageIndex = 0
	}
}
