package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	SwipeNone = Swipe(0)
	SwipeUp   = Swipe(1)
	SwipeDown = Swipe(2)
)

type Swipe int

// Input reads for example drag input like up/down (touch support for mobile)
type Input struct {
	touchIDs []ebiten.TouchID
}

func NewInput() Input {
	return Input{}
}

func (i *Input) Dir() (Swipe, bool) {
	// TODO: catch drag (mouse and touch)...
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		log.Println("Left mouse button pressed...")
	}

	i.touchIDs = inpututil.AppendJustPressedTouchIDs(i.touchIDs[:0])
	// ...

	return SwipeNone, false
}

func (i *Input) IsSwipeUp() bool {
	return false
}

func (i *Input) IsSwipeDown() bool {
	return false
}
