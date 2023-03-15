package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Input reads for example drag input like up/down (touch support for mobile)
type Input struct {
	touchIDs []ebiten.TouchID
}

func NewInput() Input {
	return Input{}
}

func (i *Input) Dir() (key ebiten.Key, ok bool) {
	// TODO: catch drag (mouse and touch)...
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// ...
	}

	i.touchIDs = inpututil.AppendJustPressedTouchIDs(i.touchIDs[:0])
	// ...

	return
}

func (i *Input) IsSwipeUp() {
}

func (i *Input) IsSwipeDown() {
}
