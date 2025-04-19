package tiles

import (
	"bytes"
	"image"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

// Tile implements the `Tiles` interface
type Tile struct {
	Sprites      map[State]*ebiten.Image
	ImageOptions *ebiten.DrawImageOptions

	data       *TilesData
	dragFn     func(tileX float64, tileY float64) (x float64, y float64)
	thrownAway bool
}

func NewTile(d *TilesData) *Tile {
	s := make(map[State]*ebiten.Image)

	// Tile
	f, _, err := image.Decode(bytes.NewReader(images.Tile))
	if err != nil {
		panic(err)
	}
	s[StateOK] = ebiten.NewImageFromImage(f)

	// TileWithCrack
	f, _, err = image.Decode(bytes.NewReader(images.TileWithCrack))
	if err != nil {
		panic(err)
	}
	s[StateCrack] = ebiten.NewImageFromImage(f)

	// Stamp Adhesive
	f, _, err = image.Decode(bytes.NewReader(images.TileWithStampAdhesive))
	if err != nil {
		panic(err)
	}
	s[StateStampAdhesive] = ebiten.NewImageFromImage(f)

	return &Tile{
		Sprites: s,
		ImageOptions: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		data:   d,
		dragFn: nil,
	}
}

func (t *Tile) Draw(screen *ebiten.Image) {
	t.ImageOptions.GeoM.Reset()
	t.ImageOptions.GeoM.Scale(*t.data.Scale, *t.data.Scale)

	if t.dragFn != nil {
		t.data.X, t.data.Y = t.dragFn(t.data.X, t.data.Y)
	}

	t.ImageOptions.GeoM.Translate(t.data.X, t.data.Y)

	screen.DrawImage(t.Sprites[t.data.State], t.ImageOptions)
}

func (t *Tile) Size() (w, h float64) {
	iW := t.Sprites[t.data.State].Bounds().Dx()
	iH := t.Sprites[t.data.State].Bounds().Dy()
	return float64(iW) * *t.data.Scale, float64(iH) * *t.data.Scale
}

func (t *Tile) Data() *TilesData {
	return t.data
}

func (t *Tile) SetDraggedFn(fn func(tileX float64, tileY float64) (x float64, y float64)) {
	t.dragFn = fn
}

func (t *Tile) ThrowAway() {
	t.thrownAway = true
}

func (t *Tile) IsThrownAway() bool {
	return t.thrownAway
}
