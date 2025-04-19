package component

import (
	"math/rand"
	"slices"
	"the-press-department/internal/sprites"
	"the-press-department/internal/stats"
	"the-press-department/internal/tiles"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// RollingRailway controls the press stop/start and moves all the
// tiles based on the engines
//
// NOTE: Create and engines type only if multiple engines are needed
type RollingRailway struct {
	input                     Component[RollingRailwayUserInputData]
	data                      *RollingRailwayData
	rolls                     []Coord
	scale                     *float64
	screenWidth, screenHeight float64
	lastUpdate                time.Time
	lastTile                  time.Time
	rand                      *rand.Rand
	tileStates                []tiles.State
}

func NewRollingRailway(scale *float64, data *RollingRailwayData) Component[RollingRailwayData] {
	c := &RollingRailway{
		input:      NewRollingRailwayUserInput(&RollingRailwayUserInputData{}),
		scale:      scale,
		data:       data,
		rolls:      make([]Coord, 0),
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
		rand:       rand.New(rand.NewSource(time.Now().Unix())),
		tileStates: []tiles.State{
			tiles.StateOK,
			tiles.StateOK,
			tiles.StateCrack,
		},
	}

	return c
}

func (c *RollingRailway) Layout(outsideWidth, outsideHeight int) (int, int) {
	c.screenWidth = float64(outsideWidth)
	c.screenHeight = float64(outsideHeight)

	c.input.Layout(outsideWidth, outsideHeight)

	if c.screenHeight != float64(outsideHeight) {
		c.screenHeight = float64(outsideHeight)

		// update tiles
		for _, t := range c.data.tiles {
			if !t.IsThrownAway() {
				_, h := t.Size()
				t.Data().Y = (c.screenHeight / 2) - (h / 2)
			}
		}
	}

	return outsideWidth, outsideHeight
}

func (c *RollingRailway) Update() error {
	c.data.X = c.data.x
	c.data.Y = c.data.y

	next := time.Now()
	c.Data().SetUpdateData(
		c.calcRange(next), // r
		0,                 // x
		c.screenHeight/2-(c.Data().Height()/2), // y
		c.screenWidth, // width
	)

	c.data.SetSprite()
	w, _ := c.data.Roll.GetAssetSize()
	padding := w * 3

	// Update roll coords (x axis)
	c.rolls = make([]Coord, 0)
	for p := c.data.X; p <= c.data.size; p += (w + padding) {
		c.rolls = append(c.rolls, Coord{X: float64(p), Y: c.data.Y})
	}

	// Only handle user input if not on Pause
	if !c.data.Pause {
		// Handle user input
		c.input.Data().ThrowAwayPaddingTop = c.Data().Y - 10
		c.input.Data().ThrowAwayPaddingBottom = c.Data().Y + c.Data().Height() + 10
		c.input.Data().Tiles = c.data.tiles
		_ = c.input.Update()
	}

	// Move tiles
	c.updateTiles(next)

	// Press a new tile
	c.updatePress(next)

	// Set the last update field
	c.lastUpdate = next

	return nil
}

func (c *RollingRailway) Draw(screen *ebiten.Image) {
	for i := range c.rolls {
		c.data.Roll.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}

	// Draw the tile with the given positions
	for _, tile := range c.data.tiles {
		tile.Draw(screen)
	}
}

func (c *RollingRailway) Data() *RollingRailwayData {
	return c.data
}

func (c *RollingRailway) updatePress(next time.Time) {
	if c.data.Pause {
		// on pause just add the diff between next and last and add it to last
		c.lastTile = c.lastTile.Add(next.Sub(c.lastTile))
		return
	}

	// check time and get a tile based on BPM
	ms := time.Microsecond * time.Duration(
		60/c.data.PressBPM()*1000000,
	)
	if c.lastTile.Add(ms).UnixMicro() <= next.UnixMicro() {
		// get a new tile here
		tile := tiles.NewTile(&tiles.TilesData{
			State: c.getRandomState(),
			Scale: c.scale,
		})

		_, h := tile.Size()
		tile.Data().X = c.screenWidth
		tile.Data().Y = (c.screenHeight / 2) - (h / 2)

		c.data.tiles = append(c.data.tiles, tile)
		c.data.Stats.TilesProduced++
		c.lastTile = next
	}
}

func (c *RollingRailway) updateTiles(next time.Time) {
	toRemove := make([]tiles.Tiles, 0)

	// Update new tiles position
	for _, t := range c.data.tiles {
		d := t.Data()
		w, h := t.Size()

		// Check if tile has thrownAway state
		if t.IsThrownAway() {
			// Animation
			pressY := (c.screenHeight / 2) - (h / 2)
			r := float64(next.Sub(c.lastUpdate).Seconds()) * (250) * (*c.scale * 10)
			if d.Y <= pressY {
				d.Y -= r
			} else {
				d.Y += r
			}
		}

		// Update x position (based on time since last update)
		d.X -= c.calcRange(next)

		// Set tiles which are out of screen to remove
		if d.X <= 0-w || d.Y <= 0-h || d.Y >= c.screenHeight { // x-axis
			// Money management
			if !t.IsThrownAway() {
				switch t.Data().State {
				case tiles.StateOK:
					c.data.Stats.AddGoodTile()
				default:
					c.data.Stats.AddBadTile()
				}
			} else {
				c.data.Stats.AddThrownAwayTile(t)
			}

			toRemove = append(toRemove, t)
		}
	}

	// Remove it
	for _, t := range toRemove {
		for i, t2 := range c.data.tiles {
			if t == t2 {
				c.data.tiles = slices.Delete(c.data.tiles, i, i+1)
				break
			}
		}
	}
}

func (c *RollingRailway) getRandomState() tiles.State {
	return c.tileStates[c.rand.Intn(len(c.tileStates))]
}

func (c *RollingRailway) calcRange(next time.Time) float64 {
	return (float64(next.Sub(c.lastUpdate).Seconds()) *
		(c.data.Stats.RollingRailwayHzMultiply * c.data.Hz())) *
		(*c.scale * 10)
}

type RollingRailwayData struct {
	Stats *stats.Game
	Roll  *sprites.Roll
	Pause bool // Pause will stop the machines :)

	// TODO: Ok, i hote this: X, x, Y, y
	X, Y float64

	// Update fields
	rSum, r, x, y, size float64

	tiles []tiles.Tiles
}

func (c *RollingRailwayData) SetUpdateData(r, x, y, size float64) {
	c.r = r
	c.x = x
	c.y = y
	c.size = size
}

func (c *RollingRailwayData) SetSprite() {
	c.rSum += c.r
	w, _ := c.Roll.GetAssetSize()
	if c.rSum >= w {
		c.Roll.NextSprite()
		c.rSum = 0
	}
}

func (c *RollingRailwayData) Height() float64 {
	_, h := c.Roll.GetAssetSize()
	return h
}

func (c *RollingRailwayData) PressBPM() float64 {
	if c.Pause {
		return 0
	}

	return c.Stats.PressBPM
}

func (c *RollingRailwayData) Hz() float64 {
	if c.Pause {
		return 0
	}

	return c.Stats.RollingRailwayHz
}

func (c *RollingRailwayData) Tiles() []tiles.Tiles {
	return c.tiles
}

// RollingRailwayUserInput reads for example drag input like up/down (touch support for mobile)
type RollingRailwayUserInput struct {
	data *RollingRailwayUserInputData

	touchIDs []ebiten.TouchID

	startY float64
	lastY  float64

	tile  tiles.Tiles
	touch map[ebiten.TouchID]struct{}
}

func NewRollingRailwayUserInput(data *RollingRailwayUserInputData) Component[RollingRailwayUserInputData] {
	return &RollingRailwayUserInput{
		data:  data,
		touch: make(map[ebiten.TouchID]struct{}),
	}
}

func (i *RollingRailwayUserInput) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (i *RollingRailwayUserInput) Draw(screen *ebiten.Image) {
}

func (i *RollingRailwayUserInput) Update() error {
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

func (i *RollingRailwayUserInput) getTile(x, _ float64, tiles []tiles.Tiles) tiles.Tiles {
	for _, tile := range tiles {
		w, _ := tile.Size()
		if x >= tile.Data().X && x <= tile.Data().X+w {
			return tile
		}
	}

	return nil
}

func (i *RollingRailwayUserInput) Data() *RollingRailwayUserInputData {
	return i.data
}

type RollingRailwayUserInputData struct {
	ThrowAwayPaddingTop    float64
	ThrowAwayPaddingBottom float64
	Tiles                  []tiles.Tiles
}
