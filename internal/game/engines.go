package game

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Engines struct {
	Game *Game
	BPM  float64 // BPM are the bumps per minute (the press speed)
	MPM  float64 // MPM are the miles per seconds (the engine speed)

	tiles      []*Tile
	lastTile   time.Time
	lastUpdate time.Time

	tilesCount int

	_tile *Tile
}

func NewEngines() *Engines {
	return &Engines{
		BPM:        6,
		MPM:        8,
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
	}
}

func (e *Engines) Draw(screen *ebiten.Image) {
	// draw the tile with the given positions
	for _, e._tile = range e.tiles {
		ebitenutil.DrawRect(
			screen,                                // dst
			float64(e.Game.ScreenWidth)-e._tile.X, // x - start right
			float64(e.Game.ScreenHeight)/2-(e._tile.Height/2), // y - center
			e._tile.Width,  // width
			e._tile.Height, // height
			e._tile.Color,  // color
		)
	}
}

func (e *Engines) Update(input *Input) (err error) {
	// update existing tile positions
	var (
		next  = time.Now()
		diff  = next.Sub(e.lastUpdate)
		index int
	)

	// move tiles
	for index, e._tile = range e.tiles {
		// update x position (based on time since last update)
		e._tile.X += float64(diff.Seconds()) * 4 * e.MPM

		if e._tile.X >= (float64(e.Game.ScreenWidth) + e._tile.Width) {
			e.tiles = e.tiles[index+1:]
			break
		}
	}

	e.lastUpdate = next

	// check time and get a tile based on BPM
	if e.lastTile.Add(time.Second*time.Duration(60/e.BPM)).UnixMicro() <= next.UnixMicro() {
		// get a new tile here
		e.tiles = append(e.tiles, NewTile(60, 120))

		e.tilesCount += 1
		log.Printf("tiles produced: %d [%v]", e.tilesCount, e.tiles)

		// and update `e.lastTile`
		e.lastTile = next
	}

	return
}
