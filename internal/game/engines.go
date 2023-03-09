package game

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Engines struct {
	Game       *Game
	Conveyor   *Conveyor
	BPM        float64 // BPM are the bumps per minute (the press speed)
	Hz         float64 // MPM are the miles per seconds (the engine speed)
	HzMultiply float64
	Scale      float64
	TilesToUse []*ebiten.Image

	tiles      []*Tile
	lastTile   time.Time
	lastUpdate time.Time

	tilesCount int

	rand *rand.Rand

	index      int
	tile       *Tile
	nextUpdate time.Time
}

func NewEngines() *Engines {
	e := &Engines{
		BPM:        6,
		Hz:         8,
		HzMultiply: 2.5,
		Scale:      DefaultScale,
		TilesToUse: []*ebiten.Image{
			ImageTile,
			ImageTileWithCrack,
		},
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
		rand:       rand.New(rand.NewSource(time.Now().Unix())),
	}

	e.Conveyor = NewConveyor(&e.Scale)

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
	for _, e.tile = range e.tiles {
		e.tile.Draw(
			screen,
			float64(e.Game.ScreenWidth)-e.tile.X, // x - start right
			float64(e.Game.ScreenHeight)/2-(e.tile.GetHeight()/2), // y - center
		)
	}
}

func (e *Engines) Update(input *Input) error {
	// update existing tile positions
	e.nextUpdate = time.Now()

	e.updateConveyor()

	// move tiles
	e.updateTiles()

	// press a tile
	e.updatePress()

	// set the last update field
	e.lastUpdate = e.nextUpdate

	return nil
}

func (e *Engines) updateConveyor() {
	// TODO: update conveyor (e.Hz based)
}

func (e *Engines) updatePress() {
	// check time and get a tile based on BPM
	if e.lastTile.Add(time.Microsecond*time.Duration(60/e.BPM*1000000)).UnixMicro() <= e.nextUpdate.UnixMicro() {
		// get a new tile here
		e.tiles = append(e.tiles, NewTile(&e.Scale, e.randomTile()))

		e.tilesCount += 1

		// and update `e.lastTile`
		e.lastTile = e.nextUpdate
	}
}

func (e *Engines) updateTiles() {
	for e.index, e.tile = range e.tiles {
		// update x position (based on time since last update)
		e.tile.X += (float64(e.nextUpdate.Sub(e.lastUpdate).Seconds()) * (e.HzMultiply * e.Hz)) * (*e.tile.Scale * 10)

		if e.tile.X >= (float64(e.Game.ScreenWidth) + e.tile.GetWidth()) {
			e.tiles = e.tiles[e.index+1:]
			break
		}
	}
}

func (e *Engines) randomTile() *ebiten.Image {
	return e.TilesToUse[e.rand.Intn(len(e.TilesToUse))]
}
