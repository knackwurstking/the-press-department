package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Engines struct {
	Game       *Game
	Conveyor   Conveyor
	BPM        float64 // BPM are the bumps per minute (the press speed)
	Hz         float64 // MPM are the miles per seconds (the engine speed)
	HzMultiply float64
	TilesToUse []*ebiten.Image

	scale float64

	tiles      []Tile
	lastTile   time.Time
	lastUpdate time.Time

	tilesCount int

	rand *rand.Rand
}

func NewEngines() Engines {
	e := Engines{
		BPM:        6,
		Hz:         8,
		HzMultiply: 2.5,
		TilesToUse: []*ebiten.Image{
			ImageTile,
			ImageTileWithCrack,
		},
		scale:      DefaultScale,
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
		rand:       rand.New(rand.NewSource(time.Now().Unix())),
	}

	e.Conveyor = NewConveyor(&e.scale)

	return e
}

func (e *Engines) Draw(screen *ebiten.Image) {
	// draw the "Conveyor"
	e.Conveyor.Draw(
		screen,
		0, float64(e.Game.ScreenHeight)/2-(e.Conveyor.GetHeight()/2),
		float64(e.Game.ScreenWidth), float64(e.Game.ScreenHeight),
	)

	// draw the tile with the given positions
	for _, tile := range e.tiles {
		tile.Draw(
			screen,
			float64(e.Game.ScreenWidth)-tile.X, // x - start right
			float64(e.Game.ScreenHeight)/2-(tile.GetHeight()/2), // y - center
		)
	}
}

func (e *Engines) Update(input Input) error {
	// update existing tile positions
	next := time.Now()

	e.updateConveyor()

	// move tiles
	e.updateTiles(next)

	// press a tile
	e.updatePress(next)

	// set the last update field
	e.lastUpdate = next

	return nil
}

func (e *Engines) updateConveyor() {
	// TODO: update conveyor (e.Hz based)
}

func (e *Engines) updatePress(next time.Time) {
	// check time and get a tile based on BPM
	if e.lastTile.Add(time.Microsecond*time.Duration(60/e.BPM*1000000)).UnixMicro() <= next.UnixMicro() {
		// get a new tile here
		e.tiles = append(e.tiles, NewTile(&e.scale, e.randomTile()))

		e.tilesCount += 1

		// and update `e.lastTile`
		e.lastTile = next
	}
}

func (e *Engines) updateTiles(next time.Time) {
	for i, t := range e.tiles {
		// update x position (based on time since last update)
		t.X += (float64(next.Sub(e.lastUpdate).Seconds()) * (e.HzMultiply * e.Hz)) * (*t.scale * 10)

		if t.X >= (float64(e.Game.ScreenWidth) + t.GetWidth()) {
			e.tiles = e.tiles[i+1:]
			break
		}
	}
}

func (e *Engines) randomTile() *ebiten.Image {
	return e.TilesToUse[e.rand.Intn(len(e.TilesToUse))]
}
