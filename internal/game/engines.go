package game

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnginesData struct {
	Stats *Stats
	Pause bool // Pause will top the machines :)

	Scale float64

	tiles []Tiles
}

func (c *EnginesData) GetTiles() []Tiles {
	return c.tiles
}

func (c *EnginesData) GetBPM() float64 {
	if c.Pause {
		return 0
	}
	return c.Stats.PressBPM
}

func (c *EnginesData) SetBPM(bpm float64) {
	c.Stats.PressBPM = bpm
}

func (c *EnginesData) GetHz() float64 {
	if c.Pause {
		return 0
	}
	return c.Stats.ConveyorHz
}

func (c *EnginesData) SetHz(n float64) {
	c.Stats.ConveyorHz = n
}

func (c *EnginesData) GetHzMultiply() float64 {
	return c.Stats.ConveyorHzMultiply
}

func (c *EnginesData) SetHzMultiply(n float64) {
	c.Stats.ConveyorHzMultiply = n
}

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Engines struct {
	conveyor Component[ConveyorData]
	input    Component[EnginesInputData]

	data                      *EnginesData
	screenWidth, screenHeight float64

	lastTile   time.Time
	lastUpdate time.Time

	rand       *rand.Rand
	tileStates []State
}

func NewEngines(data *EnginesData) *Engines {
	e := &Engines{
		input:      NewEnginesInput(&EnginesInputData{}),
		data:       data,
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
		rand:       rand.New(rand.NewSource(time.Now().Unix())),
		tileStates: []State{
			StateOK,
			StateOK,
			StateCrack,
		},
	}

	e.conveyor = NewConveyor(&ConveyorData{
		Scale:      &e.data.Scale,
		HzMultiply: e.data.GetHzMultiply(),
		Sprite:     NewRollSprite(&e.data.Scale),
	})

	return e
}

func (e *Engines) Layout(outsideWidth, outsideHeight int) (int, int) {
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

func (e *Engines) Update() error {
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

func (e *Engines) Draw(screen *ebiten.Image) {
	// Draw the "Conveyor"
	e.conveyor.Draw(screen)

	// Draw the tile with the given positions
	for _, tile := range e.data.tiles {
		tile.Draw(screen)
	}
}

func (e *Engines) Data() *EnginesData {
	return e.data
}

func (e *Engines) updateConveyor(next time.Time) {
	e.conveyor.Data().Hz = e.data.GetHz()
	e.conveyor.Data().HzMultiply = e.data.GetHzMultiply()
	e.conveyor.Data().SetUpdateData(
		e.calcR(next), // r
		0,             // x
		e.screenHeight/2-(e.conveyor.Data().GetHeight()/2), // y
		e.screenWidth, // width
	)
	_ = e.conveyor.Update()
}

func (e *Engines) updatePress(next time.Time) {
	if e.data.Pause {
		// on pause just add the diff between next and last and add it to last
		e.lastTile = e.lastTile.Add(next.Sub(e.lastTile))
		return
	}

	// check time and get a tile based on BPM
	if e.lastTile.Add(time.Microsecond*time.Duration(60/e.data.GetBPM()*1000000)).UnixMicro() <= next.UnixMicro() {
		// get a new tile here
		var tile = NewTile(&TilesData{
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

func (e *Engines) updateTiles(next time.Time) {
	toRemove := make([]Tiles, 0)

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
		d.X -= e.calcR(next)

		// Set tiles which are out of screen to remove
		if d.X <= 0-w || d.Y <= 0-h || d.Y >= e.screenHeight { // x-axis
			// Money management
			if !t.IsThrownAway() {
				switch t.Data().State {
				case StateOK:
					log.Println("Add a good tile", t)
					e.data.Stats.AddGoodTile()
				default:
					log.Println("Add a bad tile", t)
					e.data.Stats.AddBadTile()
				}
			} else {
				log.Println("Add tile throw away", t)
				e.data.Stats.AddThrownAwayTile(t)
			}

			toRemove = append(toRemove, t)
		}
	}

	// Remove it
	for _, t := range toRemove {
		// Money management
		if t.IsThrownAway() {
			// Handle game stats counter for "Money" (check the tile state first)
			log.Println("Add tile throw away", t)
			e.data.Stats.AddThrownAwayTile(t)
		}

		for i, t2 := range e.data.tiles {
			if t == t2 {
				e.data.tiles = append(e.data.tiles[:i], e.data.tiles[i+1:]...)
				break
			}
		}
	}
}

func (e *Engines) calcR(next time.Time) float64 {
	return (float64(next.Sub(e.lastUpdate).Seconds()) * (e.data.GetHzMultiply() * e.data.GetHz())) * (e.data.Scale * 10)
}

func (e *Engines) getRandomState() State {
	return e.tileStates[e.rand.Intn(len(e.tileStates))]
}
