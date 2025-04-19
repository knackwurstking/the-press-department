// TODO: Ok, rename this shit back to conveiour
package component

import (
	"math/rand"
	"slices"
	"the-press-department/internal/tiles"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// RollerConveyor controls the press stop/start and moves all the
// tiles based on the engines
//
// NOTE: Create and engines type only if multiple engines are needed
type RollerConveyor struct {
	input                     Component[RollerConveyorUserInputData]
	data                      *RollerConveyorData
	rolls                     []Coord
	scale                     *float64
	screenWidth, screenHeight float64
	lastUpdate                time.Time
	lastTile                  time.Time
	rand                      *rand.Rand
	tileStates                []tiles.State
}

func NewRollerConveyor(scale *float64, data *RollerConveyorData) Component[RollerConveyorData] {
	c := &RollerConveyor{
		input:      NewRollerConveyorUserInput(&RollerConveyorUserInputData{}),
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

func (c *RollerConveyor) Layout(outsideWidth, outsideHeight int) (int, int) {
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

func (c *RollerConveyor) Update() error {
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

func (c *RollerConveyor) Draw(screen *ebiten.Image) {
	for i := range c.rolls {
		c.data.Roll.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}

	// Draw the tile with the given positions
	for _, tile := range c.data.tiles {
		tile.Draw(screen)
	}
}

func (c *RollerConveyor) Data() *RollerConveyorData {
	return c.data
}

func (c *RollerConveyor) updatePress(next time.Time) {
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

func (c *RollerConveyor) updateTiles(next time.Time) {
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

func (c *RollerConveyor) getRandomState() tiles.State {
	return c.tileStates[c.rand.Intn(len(c.tileStates))]
}

func (c *RollerConveyor) calcRange(next time.Time) float64 {
	return (float64(next.Sub(c.lastUpdate).Seconds()) *
		(c.data.Stats.RollerConveyorHzMultiply * c.data.Hz())) *
		(*c.scale * 10)
}
