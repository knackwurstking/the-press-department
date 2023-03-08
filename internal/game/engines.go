package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Engines struct {
	BPM float64 // BPM are the bumps per minute (the press speed)
	MPM float64 // MPM are the miles per seconds (the engine speed)

	tiles      []*Tile
	lastTile   int64
	lastUpdate int64
}

func NewEngines() *Engines {
	return &Engines{
		BPM:        3.5,
		MPM:        0.5,
		lastTile:   time.Now().UnixMicro(),
		lastUpdate: time.Now().UnixMicro(),
	}
}

func (*Engines) Draw(screen *ebiten.Image) {
	// draw the tile with the given positions
}

func (e *Engines) Update(input *Input) error {
	// update existing tile positions
	var (
		next = time.Now().UnixMicro()
		diff = next - e.lastUpdate
		tile *Tile
	)

	for _, tile = range e.tiles {
		// update x position (based on time since last update)
		tile.Position.X += int64(float64(diff) * e.MPM)
	}

	// check time and get a tile (6 bumps per minute)
	if next+int64(e.BPM) >= e.lastTile+int64(e.BPM) {
		// get a new tile here
		e.tiles = append(e.tiles, NewTile(NewTilePosition(0)))

		// and update `e.lastTile`
		e.lastTile = next
	}

	return nil
}
