package tiles

import "github.com/hajimehoshi/ebiten/v2"

type Tiles interface {
	Data() *TilesData
	Draw(screen *ebiten.Image)
	Size() (w, h float64)
	ThrowAway()
	IsOK() bool
	HasStampAdhesive() bool
	HasCrack() bool
	IsThrownAway() bool
	SetDraggedFn(func(tX float64, tY float64) (x float64, y float64))
}

type TilesData struct {
	State State
	Scale *float64
	X, Y  float64
}
