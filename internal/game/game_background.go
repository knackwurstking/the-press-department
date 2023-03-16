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

type BackgroundConfig struct {
	Scale float64
	Image *ebiten.Image
}

// Background for the game (just some shit with grey)
type Background struct {
	game         *Game
	config       *BackgroundConfig
	imageOptions *ebiten.DrawImageOptions
}

func NewBackground(config *BackgroundConfig) *Background {
	return &Background{
		config: config,
		imageOptions: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
	}
}

func (b *Background) Update() error {
	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	w, h := b.config.Image.Size()
	imageWidth := float64(w) * b.config.Scale
	imageHeight := float64(h) * b.config.Scale
	col := int(math.Ceil(float64(b.game.ScreenWidth) / imageWidth))
	row := int(math.Ceil(float64(b.game.ScreenHeight) / imageHeight))

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			b.imageOptions.GeoM.Reset()
			b.imageOptions.GeoM.Scale(b.config.Scale, b.config.Scale)
			b.imageOptions.GeoM.Translate(
				imageWidth*float64(c),
				imageHeight*float64(r),
			)
			screen.DrawImage(b.config.Image, b.imageOptions)
		}
	}
}

func (b *Background) SetGame(game *Game) {
	b.game = game
}

func (b *Background) SetConfig(config *BackgroundConfig) {
	b.config = config
}

func (b *Background) GetConfig() *BackgroundConfig {
	return b.config
}
