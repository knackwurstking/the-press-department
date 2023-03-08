package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Engines struct {
	Game *Game
	BPM  float64 // BPM are the bumps per minute (the press speed)
	Hz   float64 // MPM are the miles per seconds (the engine speed)

	tiles      []*Tile
	lastTile   time.Time
	lastUpdate time.Time

	tilesCount int

	_tile       *Tile
	_nextUpdate time.Time
}

func NewEngines() *Engines {
	return &Engines{
		BPM:        6,
		Hz:         8,
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
	}
}

func (e *Engines) Draw(screen *ebiten.Image) {
	// draw the tile with the given positions
	for _, e._tile = range e.tiles {
		e._tile.Draw(
			screen,
			float64(e.Game.ScreenWidth)-e._tile.X, // x - start right
			float64(e.Game.ScreenHeight)/2-(e._tile.Height/2), // y - center
		)
	}
}

func (e *Engines) Update(input *Input) error {
	// update existing tile positions
	e._nextUpdate = time.Now()

	// move tiles
	e.updateTiles()

	// press a tile
	e.updatePress()

	// set the last update field
	e.lastUpdate = e._nextUpdate

	return nil
}

func (e *Engines) updatePress() {
	// check time and get a tile based on BPM
	if e.lastTile.Add(time.Microsecond*time.Duration(60/e.BPM*1000000)).UnixMicro() <= e._nextUpdate.UnixMicro() {
		// get a new tile here
		e.tiles = append(e.tiles, NewTile(60, 120))

		e.tilesCount += 1

		// and update `e.lastTile`
		e.lastTile = e._nextUpdate
	}
}

func (e *Engines) updateTiles() {
	var i int
	for i, e._tile = range e.tiles {
		// update x position (based on time since last update)
		e._tile.X += float64(e._nextUpdate.Sub(e.lastUpdate).Seconds()) * 3 * e.Hz

		if e._tile.X >= (float64(e.Game.ScreenWidth) + e._tile.Width) {
			e.tiles = e.tiles[i+1:]
			break
		}
	}
}
