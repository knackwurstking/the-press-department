package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type EnginesInputData struct {
	ThrowAwayPaddingTop    float64
	ThrowAwayPaddingBottom float64
	Tiles                  []Tiles
}

// Input reads for example drag input like up/down (touch support for mobile)
type EnginesInput struct {
	data *EnginesInputData

	touchIDs []ebiten.TouchID

	startY float64
	lastY  float64

	tile  Tiles
	touch map[ebiten.TouchID]struct{}
}

func NewEnginesInput(data *EnginesInputData) *EnginesInput {
	return &EnginesInput{
		data:  data,
		touch: make(map[ebiten.TouchID]struct{}),
	}
}

func (i *EnginesInput) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (i *EnginesInput) Draw(screen *ebiten.Image) {
}

func (i *EnginesInput) Update() error {
	// handle mouse input
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		i.tile = i.checkForTile(float64(x), float64(y), i.data.Tiles)
		if i.tile != nil {
			i.startY = float64(y)
			i.lastY = i.startY

			i.tile.SetDraggedFn(func(tX, tY float64) (x float64, y float64) {
				_, _y := ebiten.CursorPosition()
				tY -= i.lastY - float64(_y)
				i.lastY = float64(_y)
				return tX, tY
			})
		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if i.tile != nil {
			i.tile.SetDraggedFn(nil)

			_, h := i.tile.Size()
			if i.tile.Data().Y+h > i.data.ThrowAwayPaddingBottom ||
				i.tile.Data().Y < i.data.ThrowAwayPaddingTop {
				i.tile.ThrowAway()
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
		i.tile = i.checkForTile(float64(x), float64(y), i.data.Tiles)
		if i.tile != nil {
			i.startY = float64(y)
			i.lastY = i.startY

			i.tile.SetDraggedFn(func(tX, tY float64) (x float64, y float64) {
				_x, _y := ebiten.TouchPosition(touchID)
				if _x == 0 && _y == 0 {
					i.tile.SetDraggedFn(nil)

					_, h := i.tile.Size()
					if i.tile.Data().Y+h > i.data.ThrowAwayPaddingBottom ||
						i.tile.Data().Y < i.data.ThrowAwayPaddingTop {
						i.tile.ThrowAway()
					}

					i.tile = nil
					return tX, tY
				}

				tY -= i.lastY - float64(_y)
				i.lastY = float64(_y)
				return tX, tY
			})
		}
	}

	return nil
}

func (i *EnginesInput) checkForTile(x, y float64, tiles []Tiles) Tiles {
	for _, tile := range tiles {
		w, _ := tile.Size()
		if x >= tile.Data().X && x <= tile.Data().X+w {
			return tile
		}
	}

	return nil
}

func (i *EnginesInput) Data() *EnginesInputData {
	return i.data
}
