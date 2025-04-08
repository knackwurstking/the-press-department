package tiles

import (
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tiles interface {
	Data() *TilesData
	Draw(screen *ebiten.Image)
	Size() (w, h float64)
	ThrowAway()
	IsThrownAway() bool
	SetDraggedFn(func(tX float64, tY float64) (x float64, y float64))
}

type TilesData struct {
	State State
	Scale *float64
	X, Y  float64
}
