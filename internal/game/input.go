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
	Game *Game

	ThrowAwayPaddingTop    float64
	ThrowAwayPaddingBottom float64

	touchIDs []ebiten.TouchID

	startY float64
	lastY  float64

	tile  *Tile
	touch map[ebiten.TouchID]struct{}
}

func NewInput() *Input {
	return &Input{
		touch: make(map[ebiten.TouchID]struct{}),
	}
}

func (i *Input) Update(tiles []*Tile) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		log.Println("Left mouse button pressed...")

		x, y := ebiten.CursorPosition()

		i.tile = i.checkForTile(float64(x), float64(y), tiles)
		if i.tile != nil {
			log.Printf("Got a tile at the cursors position (%d, %d) %#v", x, y, i.tile)

			i.startY = float64(y)
			i.lastY = i.startY
			i.tile.SetDragged(func(tileX, tileY float64) (float64, float64) {
				_, _y := ebiten.CursorPosition()
				tileY -= i.lastY - float64(_y)
				i.lastY = float64(_y)
				return tileX, tileY
			})
		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		log.Println("release mouse button", i.tile)

		if i.tile != nil {
			i.tile.SetDragged(nil)

			if i.tile.Y+i.tile.GetHeight() > i.ThrowAwayPaddingBottom+(i.tile.GetHeight()/6) ||
				i.tile.Y < i.ThrowAwayPaddingTop-(i.tile.GetHeight()/6) {
				i.tile.SetThrownAway()
			}

			i.tile = nil
		}
	}

	// TODO: catch drag (touch)...
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
