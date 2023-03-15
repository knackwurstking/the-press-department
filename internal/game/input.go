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

type SwipeTile struct {
	swipeType SwipeType
	tile      *Tile
}

func (s *SwipeTile) GetType() SwipeType {
	return s.swipeType
}

func (s *SwipeTile) GetTile() *Tile {
	return s.tile
}

func NewSwipeTile(t SwipeType, tile *Tile) *SwipeTile {
	return &SwipeTile{
		swipeType: t,
		tile:      tile,
	}
}

// Input reads for example drag input like up/down (touch support for mobile)
type Input struct {
	Game     *Game
	touchIDs []ebiten.TouchID
}

func NewInput() Input {
	return Input{}
}

func (i *Input) Dir(tiles []*Tile) (*SwipeTile, bool) {
	// TODO: catch drag (mouse and touch)...
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		log.Println("Left mouse button pressed...")

		// TODO: get pressed tile and create `MouseDraggingObject` struct or
		//  	 	 return if no tile at the cursors position

		x, y := ebiten.CursorPosition()
		tile := i.checkForTile(float64(x), float64(y), tiles)
		if tile != nil {
			log.Printf("Got a tile at the cursors position (%d, %d) %#v", x, y, tile)
			return NewSwipeTile(SwipeNone, tile), true
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
			return NewSwipeTile(SwipeNone, tile), true
		}
	}

	return nil, false
}

func (i *Input) checkForTile(x, y float64, tiles []*Tile) *Tile {
	for _, tile := range tiles {
		txS := float64(i.Game.ScreenWidth) - tile.X
		txE := float64(i.Game.ScreenWidth) - tile.X + tile.GetWidth()
		if x >= txS && x <= txE {
			return tile
		}
	}

	return nil
}
