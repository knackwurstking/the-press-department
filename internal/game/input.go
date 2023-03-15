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

func (i *Input) Dir(tiles []*Tile) (Swipe, bool) {
	// TODO: catch drag (mouse and touch)...
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		log.Println("Left mouse button pressed...")

		// TODO: get pressed tile and create `MouseDraggingObject` struct or
		//  	 	 return if no tile at the cursors position

		x, y := ebiten.CursorPosition()
		tile := i.checkForTile(float64(x), float64(y), tiles)
		if tile != nil {
			log.Printf("Got a tile at the cursors position (%d, %d) %#v", x, y, tile)
		}
	}

	i.touchIDs = inpututil.AppendJustPressedTouchIDs(i.touchIDs[:0])
	for _, touchID := range i.touchIDs {
		// TODO: pass touch id to `MouseDraggingObject` struct if position is on a
		//   		 tile or leave

		x, y := ebiten.TouchPosition(touchID)
		tile := i.checkForTile(float64(x), float64(y), tiles)
		if tile != nil {
			log.Printf("Got a tile at touch position (%d, %d) %#v", x, y, tile)
		}
	}

	return SwipeNone, false
}

func (i *Input) IsSwipeUp() bool {
	return false
}

func (i *Input) IsSwipeDown() bool {
	return false
}

func (i *Input) checkForTile(x, y float64, tiles []*Tile) *Tile {
	for _, tile := range tiles {
		if tile.X >= x && x <= tile.GetWidth() {
			return tile
		}
	}

	return nil
}
