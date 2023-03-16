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
	TilesToUse []*ebiten.Image

	hz         float64 // MPM are the miles per seconds (the engine speed)
	hzMultiply float64
	scale      float64

	tiles      []*Tile
	lastTile   time.Time
	lastUpdate time.Time

	tilesCount int

	rand *rand.Rand
}

func NewEngines(scale float64) *Engines {
	e := &Engines{
		BPM: 6,
		TilesToUse: []*ebiten.Image{
			ImageTile,
			ImageTile,
			ImageTile,
			ImageTileWithCrack,
		},
		hz:         8, // NOTE: cycles/seconds
		hzMultiply: 2.5,
		scale:      scale,
		lastTile:   time.Now(),
		lastUpdate: time.Now(),
		rand:       rand.New(rand.NewSource(time.Now().Unix())),
	}

	e.Conveyor = NewConveyor(&e.scale, e.hzMultiply)

	return e
}

func (e *Engines) Draw(screen *ebiten.Image) {
	// draw the "Conveyor"
	e.Conveyor.Draw(screen)

	// draw the tile with the given positions
	for _, tile := range e.tiles {
		tile.Draw(screen)
	}
}

func (e *Engines) Update(input *Input) error {
	input.ThrowAwayPaddingTop = e.Conveyor.Y
	input.ThrowAwayPaddingBottom = e.Conveyor.Y + e.Conveyor.GetHeight()
	input.Update(e.tiles)

	// update existing tile positions
	next := time.Now()

	e.updateConveyor(next)

	// move tiles
	e.updateTiles(next)

	// press a tile
	e.updatePress(next)

	// set the last update field
	e.lastUpdate = next

	return nil
}

func (e *Engines) GetHz() float64 {
	return e.hz
}

func (e *Engines) GetHzMultiply() float64 {
	return e.hzMultiply
}

func (e *Engines) updateConveyor(next time.Time) {
	e.Conveyor.hz = e.hz
	e.Conveyor.hzMultiply = e.hzMultiply
	e.Conveyor.Update(
		e.calcR(next),
		0, // x
		float64(e.Game.ScreenHeight)/2-(e.Conveyor.GetHeight()/2), // y
		float64(e.Game.ScreenWidth),                               // width
	)
}

func (e *Engines) updatePress(next time.Time) {
	// check time and get a tile based on BPM
	if e.lastTile.Add(time.Microsecond*time.Duration(60/e.BPM*1000000)).UnixMicro() <= next.UnixMicro() {
		// get a new tile here
		tile := NewTile(e.scale, e.randomTile())

		tile.X = float64(e.Game.ScreenWidth)
		tile.Y = (float64(e.Game.ScreenHeight) / 2) - (tile.GetHeight() / 2)

		e.tiles = append(e.tiles, tile)

		e.tilesCount += 1

		// and update `e.lastTile`
		e.lastTile = next
	}
}

func (e *Engines) updateTiles(next time.Time) {
	toRemove := make([]*Tile, 0)
	for i := 0; i < len(e.tiles); i++ {
		// check for thrownAway tiles firs
		if e.tiles[i].IsThrownAway() {
			// throw away animation (up if tiles y lower then the initial y set from the press)
			pressY := (float64(e.Game.ScreenHeight) / 2) - (e.tiles[i].GetHeight() / 2)
			r := float64(next.Sub(e.lastUpdate).Seconds()) * (250) * (e.scale * 10)
			if e.tiles[i].Y <= pressY {
				e.tiles[i].Y -= r
			} else {
				e.tiles[i].Y += r
			}
		}

		// update x position (based on time since last update)
		e.tiles[i].X -= e.calcR(next)

		if e.tiles[i].X <= 0-e.tiles[i].GetWidth() {
			e.tiles = e.tiles[i+1:]
			break
		} else if e.tiles[i].Y <= 0-e.tiles[i].GetHeight() || e.tiles[i].Y >= float64(e.Game.ScreenHeight) {
			toRemove = append(toRemove, e.tiles[i])
		}
	}

	for _, t := range toRemove {
		for i, t2 := range e.tiles {
			if t == t2 {
				e.tiles = append(e.tiles[:i], e.tiles[i+1:]...)
				break
			}
		}
	}
}

func (e *Engines) calcR(next time.Time) float64 {
	return (float64(next.Sub(e.lastUpdate).Seconds()) * (e.hzMultiply * e.hz)) * (e.scale * 10)
}

func (e *Engines) randomTile() *ebiten.Image {
	return e.TilesToUse[e.rand.Intn(len(e.TilesToUse))]
}
