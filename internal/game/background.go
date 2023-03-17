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

type BackgroundData struct {
	Scale float64
	Image *ebiten.Image
}

// Background for the game (just some shit with grey)
type Background struct {
	data                      *BackgroundData
	imageOptions              *ebiten.DrawImageOptions
	screenWidth, screenHeight float64
}

func NewBackground(data *BackgroundData) *Background {
	return &Background{
		data: data,
		imageOptions: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
	}
}

func (b *Background) Layout(outsideWidth, outsideHeight int) (int, int) {
	b.screenWidth = float64(outsideWidth)
	b.screenHeight = float64(outsideHeight)

	return outsideWidth, outsideHeight
}

func (b *Background) Update() error {
	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	w, h := b.data.Image.Size()
	imageWidth := float64(w) * b.data.Scale
	imageHeight := float64(h) * b.data.Scale
	col := int(math.Ceil(b.screenWidth / imageWidth))
	row := int(math.Ceil(b.screenHeight / imageHeight))

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			b.imageOptions.GeoM.Reset()
			b.imageOptions.GeoM.Scale(b.data.Scale, b.data.Scale)
			b.imageOptions.GeoM.Translate(
				imageWidth*float64(c),
				imageHeight*float64(r),
			)
			screen.DrawImage(b.data.Image, b.imageOptions)
		}
	}
}

func (b *Background) Data() *BackgroundData {
	return b.data
}
