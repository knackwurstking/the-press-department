package component

import (
	"math/rand"
	"the-press-department/internal/stats"
	"the-press-department/internal/tiles"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Engine struct {
	// TODO: Kicked from the Engine
	conveyor Component[RollingRailwayData]
	input    Component[EngineUserInputData]

	data                      *EngineData
	screenWidth, screenHeight float64

	lastTile   time.Time
	lastUpdate time.Time

	rand       *rand.Rand
	tileStates []tiles.State
}

func NewEngine(data *EngineData) *Engine {
	e := &Engine{
		input:      NewEngineUserInput(&EngineUserInputData{}),
		data:       data,
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
		rand:       rand.New(rand.NewSource(time.Now().Unix())),
		tileStates: []tiles.State{
			tiles.StateOK,
			tiles.StateOK,
			tiles.StateCrack,
		},
	}

	e.conveyor = NewRollingRailway(&RollingRailwayData{
		Scale:      &e.data.Scale,
		HzMultiply: e.data.GetHzMultiply(),
		Sprite:     NewRollSprite(&e.data.Scale),
	})

	return e
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	if e.screenHeight != float64(outsideHeight) {
		e.screenHeight = float64(outsideHeight)

		// update tiles
		for _, t := range e.data.tiles {
			if !t.IsThrownAway() {
				_, h := t.Size()
				t.Data().Y = (e.screenHeight / 2) - (h / 2)
			}
		}
	}

	e.screenWidth = float64(outsideWidth)

	e.conveyor.Layout(outsideWidth, outsideHeight)
	e.input.Layout(outsideWidth, outsideHeight)

	return outsideWidth, outsideWidth
}

func (e *Engine) Update() error {
	next := time.Now()
	e.updateConveyor(next)

	// Press a tile
	e.updatePress(next)

	// Only handle user input if not on Pause
	if !e.data.Pause {
		// Handle user input
		e.input.Data().ThrowAwayPaddingTop = e.conveyor.Data().Y - 10
		e.input.Data().ThrowAwayPaddingBottom = e.conveyor.Data().Y + e.conveyor.Data().GetHeight() + 10
		e.input.Data().Tiles = e.data.tiles
		_ = e.input.Update()
	}

	// Move tiles
	e.updateTiles(next)

	// Set the last update field
	e.lastUpdate = next

	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	// Draw the "Conveyor"
	e.conveyor.Draw(screen)

	// Draw the tile with the given positions
	for _, tile := range e.data.tiles {
		tile.Draw(screen)
	}
}

func (e *Engine) Data() *EngineData {
	return e.data
}

func (e *Engine) updateConveyor(next time.Time) {
	e.conveyor.Data().Hz = e.data.GetHz()
	e.conveyor.Data().HzMultiply = e.data.GetHzMultiply()
	e.conveyor.Data().SetUpdateData(
		e.calcRange(next), // r
		0,                 // x
		e.screenHeight/2-(e.conveyor.Data().GetHeight()/2), // y
		e.screenWidth, // width
	)
	_ = e.conveyor.Update()
}

func (e *Engine) updatePress(next time.Time) {
	if e.data.Pause {
		// on pause just add the diff between next and last and add it to last
		e.lastTile = e.lastTile.Add(next.Sub(e.lastTile))
		return
	}

	// check time and get a tile based on BPM
	ms := time.Microsecond * time.Duration(
		60/e.data.GetBPM()*1000000,
	)
	if e.lastTile.Add(ms).UnixMicro() <= next.UnixMicro() {
		// get a new tile here
		tile := tiles.NewTile(&tiles.TilesData{
			State: e.getRandomState(),
			Scale: &e.data.Scale,
		})

		_, h := tile.Size()
		tile.Data().X = e.screenWidth
		tile.Data().Y = (e.screenHeight / 2) - (h / 2)

		e.data.tiles = append(e.data.tiles, tile)
		e.data.Stats.TilesProduced++
		e.lastTile = next
	}
}

func (e *Engine) updateTiles(next time.Time) {
	toRemove := make([]tiles.Tiles, 0)

	// Update new tiles position
	for _, t := range e.data.tiles {
		d := t.Data()
		w, h := t.Size()

		// Check if tile has thrownAway state
		if t.IsThrownAway() {
			// Animation
			pressY := (e.screenHeight / 2) - (h / 2)
			r := float64(next.Sub(e.lastUpdate).Seconds()) * (250) * (e.data.Scale * 10)
			if d.Y <= pressY {
				d.Y -= r
			} else {
				d.Y += r
			}
		}

		// Update x position (based on time since last update)
		d.X -= e.calcRange(next)

		// Set tiles which are out of screen to remove
		if d.X <= 0-w || d.Y <= 0-h || d.Y >= e.screenHeight { // x-axis
			// Money management
			if !t.IsThrownAway() {
				switch t.Data().State {
				case tiles.StateOK:
					e.data.Stats.AddGoodTile()
				default:
					e.data.Stats.AddBadTile()
				}
			} else {
				e.data.Stats.AddThrownAwayTile(t)
			}

			toRemove = append(toRemove, t)
		}
	}

	// Remove it
	for _, t := range toRemove {
		for i, t2 := range e.data.tiles {
			if t == t2 {
				e.data.tiles = append(e.data.tiles[:i], e.data.tiles[i+1:]...)
				break
			}
		}
	}
}

func (e *Engine) calcRange(next time.Time) float64 {
	return (float64(next.Sub(e.lastUpdate).Seconds()) * (e.data.GetHzMultiply() * e.data.GetHz())) * (e.data.Scale * 10)
}

func (e *Engine) getRandomState() tiles.State {
	return e.tileStates[e.rand.Intn(len(e.tileStates))]
}

type EngineData struct {
	Stats *stats.Game
	Pause bool // Pause will stop the machines :)

	Scale float64

	tiles []tiles.Tiles
}

func (c *EngineData) GetTiles() []tiles.Tiles {
	return c.tiles
}

func (c *EngineData) GetBPM() float64 {
	if c.Pause {
		return 0
	}
	return c.Stats.PressBPM
}

func (c *EngineData) SetBPM(bpm float64) {
	c.Stats.PressBPM = bpm
}

func (c *EngineData) GetHz() float64 {
	if c.Pause {
		return 0
	}
	return c.Stats.ConveyorHz
}

func (c *EngineData) SetHz(n float64) {
	c.Stats.ConveyorHz = n
}

func (c *EngineData) GetHzMultiply() float64 {
	return c.Stats.ConveyorHzMultiply
}

func (c *EngineData) SetHzMultiply(n float64) {
	c.Stats.ConveyorHzMultiply = n
}

// EngineInput reads for example drag input like up/down (touch support for mobile)
type EngineUserInput struct {
	data *EngineUserInputData

	touchIDs []ebiten.TouchID

	startY float64
	lastY  float64

	tile  tiles.Tiles
	touch map[ebiten.TouchID]struct{}
}

func NewEngineUserInput(data *EngineUserInputData) Component[EngineUserInputData] {
	return &EngineUserInput{
		data:  data,
		touch: make(map[ebiten.TouchID]struct{}),
	}
}

func (i *EngineUserInput) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (i *EngineUserInput) Draw(screen *ebiten.Image) {
}

func (i *EngineUserInput) Update() error {
	// handle mouse input
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		i.tile = i.getTile(float64(x), float64(y), i.data.Tiles)
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
		i.tile = i.getTile(float64(x), float64(y), i.data.Tiles)
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

func (i *EngineUserInput) getTile(x, _ float64, tiles []tiles.Tiles) tiles.Tiles {
	for _, tile := range tiles {
		w, _ := tile.Size()
		if x >= tile.Data().X && x <= tile.Data().X+w {
			return tile
		}
	}

	return nil
}

func (i *EngineUserInput) Data() *EngineUserInputData {
	return i.data
}

type EngineUserInputData struct {
	ThrowAwayPaddingTop    float64
	ThrowAwayPaddingBottom float64
	Tiles                  []tiles.Tiles
}
