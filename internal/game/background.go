package game

import (
	"bytes"
	"image"
	"math"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ImageGround image.Image
)

func init() {
	var err error

	// Ground
	ImageGround, _, err = image.Decode(bytes.NewReader(images.Ground))
	if err != nil {
		panic(err)
	}
}

// Background for the game (just some shit with grey)
type Background struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions

	scale float64
}

func NewBackground(scale float64, ground image.Image) Background {
	return Background{
		Image: ebiten.NewImageFromImage(ground),
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		scale: scale,
	}
}

func (b *Background) Draw(screen *ebiten.Image, screenWidth, screenHeight float64) {
	w, h := b.Image.Size()
	imageWidth := float64(w) * b.scale
	imageHeight := float64(h) * b.scale
	col := int(math.Ceil(screenWidth / imageWidth))
	row := int(math.Ceil(screenHeight / imageHeight))

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			b.Options.GeoM.Reset()
			b.Options.GeoM.Scale(b.scale, b.scale)
			b.Options.GeoM.Translate(
				imageWidth*float64(c),
				imageHeight*float64(r),
			)
			screen.DrawImage(b.Image, b.Options)
		}
	}
}
