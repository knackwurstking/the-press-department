package game

import (
	"bytes"
	"image"
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ImageTileAssets map[State]*ebiten.Image = make(map[State]*ebiten.Image)
)

func init() {
	// Tile
	img, _, err := image.Decode(bytes.NewReader(images.Tile))
	if err != nil {
		panic(err)
	}
	ImageTileAssets[StateOK] = ebiten.NewImageFromImage(img)

	// TileWithCrack
	img, _, err = image.Decode(bytes.NewReader(images.TileWithCrack))
	if err != nil {
		panic(err)
	}
	ImageTileAssets[StateCrack] = ebiten.NewImageFromImage(img)
}

type Tile struct {
	ImageOptions *ebiten.DrawImageOptions

	data       *TilesData
	dragFn     func(tileX float64, tileY float64) (x float64, y float64)
	thrownAway bool
}

func NewTile(d *TilesData) *Tile {
	return &Tile{
		data: d,
		ImageOptions: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
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

	screen.DrawImage(ImageTileAssets[t.data.State], t.ImageOptions)
}

func (t *Tile) Size() (w, h float64) {
	_w, _h := ImageTileAssets[t.data.State].Size()
	return float64(_w) * *t.data.Scale, float64(_h) * *t.data.Scale
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
