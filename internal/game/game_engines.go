package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnginesConfig struct {
	Pause bool // Pause will top the machines :)

	Scale float64
	Input GameComponent[EnginesInputConfig]

	bpm float64 // BPM are the bumps per minute (the press speed)

	hz         float64 // MPM are the miles per seconds (the engine speed)
	HzMultiply float64

	tilesCount int
	tiles      []*Tile
}

func (c *EnginesConfig) GetTilesCount() int {
	return c.tilesCount
}

func (c *EnginesConfig) GetTiles() []*Tile {
	return c.tiles
}

func (c *EnginesConfig) GetBPM() float64 {
	if c.Pause {
		return 0
	}
	return c.bpm
}

func (c *EnginesConfig) SetBPM(bpm float64) {
	c.bpm = bpm
}

func (c *EnginesConfig) GetHz() float64 {
	if c.Pause {
		return 0
	}
	return c.hz
}

func (c *EnginesConfig) SetHz(hz float64) {
	c.hz = hz
}

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Engines struct {
	Conveyor GameComponent[ConveyorConfig]

	config                    *EnginesConfig
	screenWidth, screenHeight float64

	tilesToUse []*ebiten.Image
	lastTile   time.Time
	lastUpdate time.Time

	rand *rand.Rand
}

func NewEngines(config *EnginesConfig) *Engines {
	e := &Engines{
		config: config,
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

	e.Conveyor = NewConveyor(&ConveyorConfig{
		Scale:      &e.config.Scale,
		HzMultiply: e.config.HzMultiply,
		Sprite:     NewRollSprite(&e.config.Scale),
	})

	return e
}

func (e *Engines) Layout(outsideWidth, outsideHeight int) (int, int) {
	if e.screenHeight != float64(outsideHeight) {
		e.screenHeight = float64(outsideHeight)

		// update tiles
		for _, t := range e.config.tiles {
			if !t.IsThrownAway() {
				t.Y = (e.screenHeight / 2) - (t.GetHeight() / 2)
			}
		}
	}

	e.screenWidth = float64(outsideWidth)

	e.Conveyor.Layout(outsideWidth, outsideHeight)
	e.config.Input.Layout(outsideWidth, outsideHeight)

	return outsideWidth, outsideWidth
}

func (e *Engines) Update() error {
	next := time.Now()
	e.updateConveyor(next)

	// Press a tile
	e.updatePress(next)

	// Handle user input
	e.config.Input.GetConfig().ThrowAwayPaddingTop = e.Conveyor.GetConfig().Y - 10
	e.config.Input.GetConfig().ThrowAwayPaddingBottom = e.Conveyor.GetConfig().Y + e.Conveyor.GetConfig().GetHeight() + 10
	e.config.Input.GetConfig().Tiles = e.config.tiles
	_ = e.config.Input.Update()

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
	for _, tile := range e.config.tiles {
		tile.Draw(screen)
	}
}

func (e *Engines) SetConfig(config *EnginesConfig) {
	e.config = config
}

func (e *Engines) GetConfig() *EnginesConfig {
	return e.config
}

func (e *Engines) updateConveyor(next time.Time) {
	e.Conveyor.GetConfig().Hz = e.config.GetHz()
	e.Conveyor.GetConfig().HzMultiply = e.config.HzMultiply
	e.Conveyor.GetConfig().SetUpdateData(
		e.calcR(next), // r
		0,             // x
		e.screenHeight/2-(e.Conveyor.GetConfig().GetHeight()/2), // y
		e.screenWidth, // width
	)
	_ = e.Conveyor.Update()
}

func (e *Engines) updatePress(next time.Time) {
	if e.config.Pause {
		// on pause just add the diff between next and last and add it to last
		e.lastTile = e.lastTile.Add(next.Sub(e.lastTile))
		return
	}

	// check time and get a tile based on BPM
	if e.lastTile.Add(time.Microsecond*time.Duration(60/e.config.GetBPM()*1000000)).UnixMicro() <= next.UnixMicro() {
		// get a new tile here
		tile := NewTile(e.config.Scale, e.randomTile())
		tile.X = e.screenWidth
		tile.Y = (e.screenHeight / 2) - (tile.GetHeight() / 2)

		e.config.tiles = append(e.config.tiles, tile)
		e.config.tilesCount += 1
		e.lastTile = next
	}
}

func (e *Engines) updateTiles(next time.Time) {
	toRemove := make([]*Tile, 0)
	for i := 0; i < len(e.config.tiles); i++ {
		// check for thrownAway tiles firs
		if e.config.tiles[i].IsThrownAway() {
			// throw away animation (up if tiles y lower then the initial y set from the press)
			pressY := (e.screenHeight / 2) - (e.config.tiles[i].GetHeight() / 2)
			r := float64(next.Sub(e.lastUpdate).Seconds()) * (250) * (e.config.Scale * 10)
			if e.config.tiles[i].Y <= pressY {
				e.config.tiles[i].Y -= r
			} else {
				e.config.tiles[i].Y += r
			}
		}

		// update x position (based on time since last update)
		e.config.tiles[i].X -= e.calcR(next)

		if e.config.tiles[i].X <= 0-e.config.tiles[i].GetWidth() {
			e.config.tiles = e.config.tiles[i+1:]
			break
		} else if e.config.tiles[i].Y <= 0-e.config.tiles[i].GetHeight() || e.config.tiles[i].Y >= e.screenHeight {
			toRemove = append(toRemove, e.config.tiles[i])
		}
	}

	for _, t := range toRemove {
		for i, t2 := range e.config.tiles {
			if t == t2 {
				e.config.tiles = append(e.config.tiles[:i], e.config.tiles[i+1:]...)
				break
			}
		}
	}
}

func (e *Engines) calcR(next time.Time) float64 {
	return (float64(next.Sub(e.lastUpdate).Seconds()) * (e.config.HzMultiply * e.config.GetHz())) * (e.config.Scale * 10)
}

func (e *Engines) randomTile() *ebiten.Image {
	return e.tilesToUse[e.rand.Intn(len(e.tilesToUse))]
}
