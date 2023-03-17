package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnginesData struct {
	Stats *Stats
	Pause bool // Pause will top the machines :)

	Scale float64
	Input Component[EnginesInputData]

	tiles []*Tile
}

func (c *EnginesData) GetTiles() []*Tile {
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
	Conveyor Component[ConveyorData]

	data                      *EnginesData
	screenWidth, screenHeight float64

	tilesToUse []*ebiten.Image
	lastTile   time.Time
	lastUpdate time.Time

	rand *rand.Rand
}

func NewEngines(data *EnginesData) *Engines {
	e := &Engines{
		data: data,
		tilesToUse: []*ebiten.Image{
			ImageTile,
			ImageTile,
			ImageTile,
			ImageTileWithCrack,
		},
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
		rand:       rand.New(rand.NewSource(time.Now().Unix())),
	}

	e.Conveyor = NewConveyor(&ConveyorData{
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
				t.Y = (e.screenHeight / 2) - (t.GetHeight() / 2)
			}
		}
	}

	e.screenWidth = float64(outsideWidth)

	e.Conveyor.Layout(outsideWidth, outsideHeight)
	e.data.Input.Layout(outsideWidth, outsideHeight)

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
		e.data.Input.GetData().ThrowAwayPaddingTop = e.Conveyor.GetData().Y - 10
		e.data.Input.GetData().ThrowAwayPaddingBottom = e.Conveyor.GetData().Y + e.Conveyor.GetData().GetHeight() + 10
		e.data.Input.GetData().Tiles = e.data.tiles
		_ = e.data.Input.Update()
	}

	// Move tiles
	e.updateTiles(next)

	// Set the last update field
	e.lastUpdate = next

	return nil
}

func (e *Engines) Draw(screen *ebiten.Image) {
	// Draw the "Conveyor"
	e.Conveyor.Draw(screen)

	// Draw the tile with the given positions
	for _, tile := range e.data.tiles {
		tile.Draw(screen)
	}
}

func (e *Engines) SetData(data *EnginesData) {
	e.data = data
}

func (e *Engines) GetData() *EnginesData {
	return e.data
}

func (e *Engines) updateConveyor(next time.Time) {
	e.Conveyor.GetData().Hz = e.data.GetHz()
	e.Conveyor.GetData().HzMultiply = e.data.GetHzMultiply()
	e.Conveyor.GetData().SetUpdateData(
		e.calcR(next), // r
		0,             // x
		e.screenHeight/2-(e.Conveyor.GetData().GetHeight()/2), // y
		e.screenWidth, // width
	)
	_ = e.Conveyor.Update()
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
		tile := NewTile(e.data.Scale, e.randomTile())
		tile.X = e.screenWidth
		tile.Y = (e.screenHeight / 2) - (tile.GetHeight() / 2)

		e.data.tiles = append(e.data.tiles, tile)
		e.data.Stats.TilesProduced++
		e.lastTile = next
	}
}

func (e *Engines) updateTiles(next time.Time) {
	toRemove := make([]*Tile, 0)

	// Update new tiles position
	for i := 0; i < len(e.data.tiles); i++ {
		// Check if tile has thrownAway state
		if e.data.tiles[i].IsThrownAway() {
			// TODO: Handle game stats counter for "Money"

			// Animation
			pressY := (e.screenHeight / 2) - (e.data.tiles[i].GetHeight() / 2)
			r := float64(next.Sub(e.lastUpdate).Seconds()) * (250) * (e.data.Scale * 10)
			if e.data.tiles[i].Y <= pressY {
				e.data.tiles[i].Y -= r
			} else {
				e.data.tiles[i].Y += r
			}
		}

		// Update x position (based on time since last update)
		e.data.tiles[i].X -= e.calcR(next)

		// Remove tiles which are out of screen
		if e.data.tiles[i].X <= 0-e.data.tiles[i].GetWidth() { // x-axis
			// TODO: Handle game stats for "Money", "GoodTiles" and "BadTiles"
			e.data.tiles = e.data.tiles[i+1:]
			break
		} else if e.data.tiles[i].Y <= 0-e.data.tiles[i].GetHeight() || e.data.tiles[i].Y >= e.screenHeight {
			toRemove = append(toRemove, e.data.tiles[i])
		}
	}

	for _, t := range toRemove {
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

func (e *Engines) randomTile() *ebiten.Image {
	return e.tilesToUse[e.rand.Intn(len(e.tilesToUse))]
}
