package game

import "github.com/hajimehoshi/ebiten/v2"

// Input reads for example drag input like up/down (touch support for mobile)
type Input struct{}

func NewInput() *Input {
	return &Input{}
}

func (*Input) Dir() (key ebiten.Key, ok bool) {
	// TODO: read input here... (Up/Down, start/end mouse drag up/down, start/end touch/swipe up/down)
	// ...

	// no valid input
	return
}
