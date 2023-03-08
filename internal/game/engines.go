package game

import (
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
	lastTile   int64
	lastUpdate int64

	_tile *Tile
}

func NewEngines() *Engines {
	return &Engines{
		BPM:        3.5,
		MPM:        0.5,
		lastTile:   time.Now().UnixMicro(),
		lastUpdate: time.Now().UnixMicro(),
	}
}

func (e *Engines) Draw(screen *ebiten.Image) {
	// draw the tile with the given positions
	for _, e._tile = range e.tiles {
		ebitenutil.DrawRect(
			screen,                         // dst
			float64(e.Game.ScreenWidth),    // x - start right
			float64(e.Game.ScreenHeight)/2, // y - center
			e._tile.Width,                  // width
			e._tile.Height,                 // height
			e._tile.Color,                  // color
		)
	}
}

func (e *Engines) Update(input *Input) (err error) {
	// update existing tile positions
	var (
		next = time.Now().UnixMicro()
		diff = next - e.lastUpdate
	)

	for _, e._tile = range e.tiles {
		// update x position (based on time since last update)
		e._tile.X += float64(diff) * e.MPM
	}

	// check time and get a tile (6 bumps per minute)
	if next+int64(e.BPM) >= e.lastTile+int64(e.BPM) {
		// get a new tile here
		e.tiles = append(e.tiles, NewTile(120, 60))

		// and update `e.lastTile`
		e.lastTile = next
	}

	return
}
