package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	SwipeNone = SwipeType(0)
	SwipeUp   = SwipeType(1)
	SwipeDown = SwipeType(2)
)

type SwipeType int

// Input reads for example drag input like up/down (touch support for mobile)
type Input struct {
	Game     *Game
	touchIDs []ebiten.TouchID

	sY        float64
	tile      *Tile
	mouseLeft bool
	touch     map[ebiten.TouchID]struct{}
}

func NewInput() Input {
	return Input{
		touch: make(map[ebiten.TouchID]struct{}),
	}
}

func (i *Input) Dir(tiles []*Tile) bool {
	// TODO: catch drag (mouse and touch)...

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		log.Println("Left mouse button pressed...")

		x, y := ebiten.CursorPosition()
		tile := i.checkForTile(float64(x), float64(y), tiles)
		if tile != nil {
			log.Printf("Got a tile at the cursors position (%d, %d) %#v", x, y, tile)
			i.mouseLeft = true
			i.tile = tile
		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// TODO: remove tile if swiped far enough
		i.mouseLeft = false
		i.tile = nil
	} else if i.mouseLeft && i.tile != nil {
		_, y := ebiten.CursorPosition()
		i.tile.Y -= i.sY - float64(y)
	}

	i.touchIDs = inpututil.AppendJustPressedTouchIDs(i.touchIDs[:0])
	for _, touchID := range i.touchIDs {
		x, y := ebiten.TouchPosition(touchID)
		tile := i.checkForTile(float64(x), float64(y), tiles)
		if tile != nil {
			log.Printf("Got a tile at touch position (%d, %d) %#v", x, y, tile)
		}
	}

	return false
}

func (i *Input) checkForTile(x, y float64, tiles []*Tile) *Tile {
	for _, tile := range tiles {
		if x >= tile.X && x <= tile.X+tile.GetWidth() {
			return tile
		}
	}

	return nil
}
