package tiles

import (
	"the-press-department/internal/images"

	"github.com/hajimehoshi/ebiten/v2"
)

var tileSet = map[State]*ebiten.Image{}

func init() {
	tileSet = map[State]*ebiten.Image{
		StateOK:            images.TileStateOK,
		StateCrack:         images.TileStateCrack,
		StateStampAdhesive: images.TileStateStampAdhesive,
	}
}

// Tile implements the `Tiles` interface
type Tile struct {
	TileSet      map[State]*ebiten.Image
	ImageOptions *ebiten.DrawImageOptions

	data       *TilesData
	dragFn     func(tileX float64, tileY float64) (x float64, y float64)
	thrownAway bool
}

func NewTile(d *TilesData) *Tile {
	return &Tile{
		TileSet: tileSet,
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

	screen.DrawImage(t.TileSet[t.data.State], t.ImageOptions)
}

func (t *Tile) Size() (w, h float64) {
	iW := t.TileSet[t.data.State].Bounds().Dx()
	iH := t.TileSet[t.data.State].Bounds().Dy()
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

func (t *Tile) IsOK() bool {
	return t.Data().State == StateOK
}

func (t *Tile) HasStampAdhesive() bool {
	return t.Data().State == StateStampAdhesive
}

func (t *Tile) HasCrack() bool {
	return t.Data().State == StateCrack
}
