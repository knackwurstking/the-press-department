package component

import (
	"math"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

type BackgroundData struct {
	image        *ebiten.Image
	imageOptions *ebiten.DrawImageOptions
}

// Background implements the `Component` interface.
// Currently it is just as background for the game (some shit with grey)
type Background struct {
	BackgroundData

	scale         *float64
	width, height float64
}

func NewBackground(scale *float64) Component[BackgroundData] {
	return &Background{
		BackgroundData: BackgroundData{
			image: images.Ground,
			imageOptions: &ebiten.DrawImageOptions{
				GeoM: ebiten.GeoM{},
			},
		},
		scale: scale,
	}
}

func (b *Background) Layout(outsideWidth, outsideHeight int) (int, int) {
	b.width = float64(outsideWidth)
	b.height = float64(outsideHeight)

	return outsideWidth, outsideHeight
}

func (b *Background) Update() error {
	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	w := b.image.Bounds().Dx()
	h := b.image.Bounds().Dy()

	imageWidth := float64(w) * *b.scale
	imageHeight := float64(h) * *b.scale

	col := int(math.Ceil(b.width / imageWidth))
	row := int(math.Ceil(b.height / imageHeight))

	for r := range row {
		for c := range col {
			b.imageOptions.GeoM.Reset()
			b.imageOptions.GeoM.Scale(*b.scale, *b.scale)
			b.imageOptions.GeoM.Translate(
				imageWidth*float64(c),
				imageHeight*float64(r),
			)
			screen.DrawImage(b.image, b.imageOptions)
		}
	}
}

func (b *Background) Data() *BackgroundData {
	return &b.BackgroundData
}
