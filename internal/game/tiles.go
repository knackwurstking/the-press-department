package game

import (
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	StateOK    = State(0)
	StateCrack = State(1)
)

type State int

type TilesData struct {
	State State
	Scale *float64
	X, Y  float64
}

type Tiles interface {
	Data() *TilesData
	Draw(screen *ebiten.Image)
	Size() (w, h float64)
	ThrowAway()
	IsThrownAway() bool
	SetDraggedFn(func(tX float64, tY float64) (x float64, y float64))
}
