package game

import (
	_ "image/jpeg"

	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"the-press-department/internal/images"
)

var (
	ImageTile          *ebiten.Image
	ImageTileWithCrack *ebiten.Image
)

func init() {
	// Tile
	img, _, err := image.Decode(bytes.NewReader(images.Tile))
	if err != nil {
		panic(err)
	}
	ImageTile = ebiten.NewImageFromImage(img)

	// TileWithCrack
	img, _, err = image.Decode(bytes.NewReader(images.TileWithCrack))
	if err != nil {
		panic(err)
	}
	ImageTileWithCrack = ebiten.NewImageFromImage(img)
}

type Tile struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions

	X float64
	Y float64

	scale float64

	dragFn     func(tileX float64, tileY float64) (x float64, y float64)
	thrownAway bool
}

func NewTile(scale float64, tile *ebiten.Image) *Tile {
	return &Tile{
		Image: tile,
		Options: &ebiten.DrawImageOptions{
			GeoM: ebiten.GeoM{},
		},
		X:      0,
		scale:  scale,
		dragFn: nil,
	}
}

func (t *Tile) Draw(screen *ebiten.Image) {
	t.Options.GeoM.Reset()
	t.Options.GeoM.Scale(t.scale, t.scale)

	if t.dragFn != nil {
		t.X, t.Y = t.dragFn(t.X, t.Y)
	}

	t.Options.GeoM.Translate(t.X, t.Y)

	screen.DrawImage(t.Image, t.Options)
}

func (t *Tile) GetHeight() float64 {
	_, h := t.Image.Size()
	return float64(h) * t.scale
}

func (t *Tile) GetWidth() float64 {
	w, _ := t.Image.Size()
	return float64(w) * t.scale
}

func (t *Tile) SetDragged(fn func(tileX float64, tileY float64) (x float64, y float64)) {
	t.dragFn = fn
}

func (t *Tile) SetThrownAway() {
	t.thrownAway = true
}

func (t *Tile) IsThrownAway() bool {
	return t.thrownAway
}

// TODO: add type State (StateCrack, StateOK)
// TODO: add public Tile field "State"
