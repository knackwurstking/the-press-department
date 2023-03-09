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

	Scale float64

	w           int
	h           int
	imageWidth  float64
	imageHeight float64
	col         int
	row         int
	r           int
	c           int
}

func NewBackground(scale float64, ground image.Image) *Background {
	return &Background{
		Image: ebiten.NewImageFromImage(ground),
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		Scale: scale,
	}
}

func (b *Background) Draw(screen *ebiten.Image, screenWidth, screenHeight float64) {
	b.w, b.h = b.Image.Size()
	b.imageWidth = float64(b.w) * b.Scale
	b.imageHeight = float64(b.h) * b.Scale

	b.col = int(math.Ceil(screenWidth / b.imageWidth))
	b.row = int(math.Ceil(screenHeight / b.imageHeight))

	for b.r = 0; b.r < b.row; b.r++ {
		for b.c = 0; b.c < b.col; b.c++ {
			b.Options.GeoM.Reset()
			b.Options.GeoM.Scale(b.Scale, b.Scale)
			b.Options.GeoM.Translate(
				b.imageWidth*float64(b.c),
				b.imageHeight*float64(b.r),
			)
			screen.DrawImage(b.Image, b.Options)
		}
	}
}
