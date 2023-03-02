package game

import (
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

func (*Board) Update(input *Input) error {
	return nil
}
