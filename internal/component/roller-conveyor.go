package component

import (
	"math/rand"
	"slices"
	"the-press-department/internal/stats"
	"the-press-department/internal/tiles"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// RollerConveyor controls the press stop/start and moves all the
// tiles based on the engines
//
// NOTE: Create and engines type only if multiple engines are needed
type RollerConveyor struct {
	RollerConveyorData

	input Component[RollerConveyorUserInputData]
	rolls []Coord

	stats *stats.Game
	scale *float64

	width, height float64
	lastUpdate    time.Time
	rand          *rand.Rand

	lastTile   time.Time
	tileStates []tiles.State
}

func NewRollerConveyor(stats *stats.Game, scale *float64, data RollerConveyorData) Component[RollerConveyorData] {
	c := &RollerConveyor{
		RollerConveyorData: data,
		input:              NewRollerConveyorUserInput(&RollerConveyorUserInputData{}),
		stats:              stats,
		rolls:              make([]Coord, 0),
		scale:              scale,
		lastTile:           time.Now(),
		lastUpdate:         time.Now(),
		rand:               rand.New(rand.NewSource(time.Now().Unix())),
		tileStates: []tiles.State{
			tiles.StateOK,
			tiles.StateOK,
			tiles.StateCrack,
			tiles.StateStampAdhesive,
			tiles.StateOK,
			tiles.StateOK,
		},
	}

	return c
}

func (c *RollerConveyor) Layout(outsideWidth, outsideHeight int) (int, int) {
	c.width = float64(outsideWidth)

	// Update tiles only if height has changed (no special reason for this)
	if c.height != float64(outsideHeight) {
		c.height = float64(outsideHeight)

		for _, t := range c.tiles {
			if !t.IsThrownAway() {
				_, h := t.Size()
				t.Data().Y = (c.height / 2) - (h / 2)
			}
		}
	}

	c.input.Layout(outsideWidth, outsideHeight)

	return outsideWidth, outsideHeight
}

func (c *RollerConveyor) Update() error {
	c.X = c.x
	c.Y = c.y

	next := time.Now()
	c.Data().SetUpdateData(
		c.calcRange(next),                // r
		0,                                // x
		c.height/2-(c.Data().Height()/2), // y
		c.width,                          // width
	)

	c.SetSprite()
	w, _ := c.RollSprite.GetAssetSize()
	padding := w * 3

	// Update roll coords (x axis)
	c.rolls = make([]Coord, 0)
	for p := c.X; p <= c.size; p += (w + padding) {
		c.rolls = append(c.rolls, Coord{X: float64(p), Y: c.Y})
	}

	// Only handle user input if not on Pause
	if !c.stats.Pause {
		// Handle user input
		c.input.Data().ThrowAwayPaddingTop = c.Data().Y - 10
		c.input.Data().ThrowAwayPaddingBottom = c.Data().Y + c.Data().Height() + 10
		c.input.Data().Tiles = c.tiles
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

func (c *RollerConveyor) Draw(screen *ebiten.Image) {
	for i := range c.rolls {
		c.RollSprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}

	// Draw the tile with the given positions
	for _, tile := range c.tiles {
		tile.Draw(screen)
	}
}

func (c *RollerConveyor) Data() *RollerConveyorData {
	return &c.RollerConveyorData
}

func (c *RollerConveyor) updatePress(next time.Time) {
	if c.stats.Pause {
		// on pause just add the diff between next and last and add it to last
		c.lastTile = c.lastTile.Add(next.Sub(c.lastTile))
		return
	}

	// check time and get a tile based on BPM
	ms := time.Microsecond * time.Duration(
		60/c.stats.PressBPM*1000000,
	)
	if c.lastTile.Add(ms).UnixMicro() <= next.UnixMicro() {
		// get a new tile here
		tile := tiles.NewTile(&tiles.TilesData{
			State: c.getRandomState(),
			Scale: c.scale,
		})

		_, h := tile.Size()
		tile.Data().X = c.width
		tile.Data().Y = (c.height / 2) - (h / 2)

		c.tiles = append(c.tiles, tile)
		c.stats.TilesProduced++
		c.lastTile = next
	}
}

func (c *RollerConveyor) updateTiles(next time.Time) {
	toRemove := make([]tiles.Tiles, 0)

	// Update new tiles position
	for _, t := range c.tiles {
		d := t.Data()
		w, h := t.Size()

		// update tiles if not thrown away
		if t.IsThrownAway() {
			pressY := (c.height / 2) - (h / 2)
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
		if d.X <= 0-w || d.Y <= 0-h || d.Y >= c.height { // x-axis
			// Money management
			if !t.IsThrownAway() {
				switch t.Data().State {
				case tiles.StateOK:
					c.stats.AddGoodTile(t)
				default:
					c.stats.AddBadTile(t)
				}
			} else {
				c.stats.AddThrownAwayTile(t)
			}

			toRemove = append(toRemove, t)
		}
	}

	// Remove it
	for _, t := range toRemove {
		for i, t2 := range c.tiles {
			if t == t2 {
				c.tiles = slices.Delete(c.tiles, i, i+1)
				break
			}
		}
	}
}

func (c *RollerConveyor) getRandomState() tiles.State {
	return c.tileStates[c.rand.Intn(len(c.tileStates))]
}

func (c *RollerConveyor) calcRange(next time.Time) float64 {
	return (float64(next.Sub(c.lastUpdate).Seconds()) *
		(c.stats.RollerConveyorHzMultiply * c.stats.RollerConveyorHz)) *
		(*c.scale * 10)
}
