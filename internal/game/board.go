package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"the-press-department/internal/game/engines"
	"the-press-department/internal/game/press"
)

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Board struct {
	Press   *press.Press
	Engines *engines.Engines
}

func NewBoard() *Board {
	return &Board{
		Press:   press.NewPress(),
		Engines: engines.NewEngines(),
	}
}

func (*Board) Draw(screen *ebiten.Image) {
	// TODO: draw the engines and the tiles from the Press
	// 			 (the Board will handle the Press and Engines)...
	// ...
}

func (*Board) Update(input *Input) error {
	return nil
}
