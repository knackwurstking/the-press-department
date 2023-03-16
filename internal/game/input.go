package game

import (
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
	// handle mouse input
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i.tile = i.checkForTile(float64(x), float64(y), tiles)
		if i.tile != nil {
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
		if i.tile != nil {
			i.tile.SetDragged(nil)

			if i.tile.Y+i.tile.GetHeight() > i.ThrowAwayPaddingBottom ||
				i.tile.Y < i.ThrowAwayPaddingTop {
				i.tile.SetThrownAway()
			}

			i.tile = nil
		}
	}

	// Handle touch input
	i.touchIDs = inpututil.AppendJustPressedTouchIDs(i.touchIDs[:0])
	if len(i.touchIDs) > 0 {
		// single finger touch
		touchID := i.touchIDs[0]
		x, y := ebiten.TouchPosition(touchID)
		i.tile = i.checkForTile(float64(x), float64(y), tiles)
		if i.tile != nil {
			i.startY = float64(y)
			i.lastY = i.startY

			i.tile.SetDragged(func(tileX, tileY float64) (float64, float64) {
				_x, _y := ebiten.TouchPosition(touchID)
				if _x == 0 && _y == 0 {
					i.tile.SetDragged(nil)

					if i.tile.Y+i.tile.GetHeight() > i.ThrowAwayPaddingBottom ||
						i.tile.Y < i.ThrowAwayPaddingTop {
						i.tile.SetThrownAway()
					}

					i.tile = nil
					return tileX, tileY
				}

				tileY -= i.lastY - float64(_y)
				i.lastY = float64(_y)
				return tileX, tileY
			})
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
