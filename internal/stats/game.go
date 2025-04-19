package stats

import (
	"the-press-department/internal/tiles"
)

// Stats for saving the game state
type Game struct {
	// Money holds the number of your money
	Money int `json:"money"`

	// GoodTiles shows the number of good tiles passed
	GoodTiles int `json:"good-tiles"`

	// BadTiles shows the number of bad tiles passed
	BadTiles int `json:"bad-tiles"`

	// TilesProduced is the number of tiles produced from press
	TilesProduced int `json:"tiles-produced"`

	// PressBPM (setup value)
	PressBPM float64 `json:"press-BPM"`

	// RollerConveyorHz (setup value)
	RollerConveyorHz float64 `json:"conveyor-hz"`

	// RollerConveyorHzMultiply (setup value)
	RollerConveyorHzMultiply float64 `json:"conveyor-hz-multiply"`

	// Pause shows if the game is running or not
	Pause bool `json:"-"`
}

func (g *Game) AddGoodTile(tile tiles.Tiles) {
	g.GoodTiles++

	if tile.IsOK() {
		g.Money += 150
	} else {
		// NOTE: Should never happen
		g.Money -= 10000
	}
}

func (g *Game) AddBadTile(tile tiles.Tiles) {
	g.BadTiles++

	if tile.HasCrack() {
		g.Money -= 400
	} else if tile.HasStampAdhesive() {
		g.Money -= 150
	} else {
		// NOTE: Should never happen, as long the tile is not ok
		g.Money -= 1000
	}
}

// "add thrown away good tile", "add thrown away bad tile"
func (g *Game) AddThrownAwayTile(tile tiles.Tiles) {
	if !tile.IsThrownAway() {
		return
	}

	// check if tile was ok (if ok then add a penalty "-1000$")...
	switch tile.Data().State {
	case tiles.StateCrack:
		g.Money -= 50
	case tiles.StateStampAdhesive:
		g.Money -= 50
	case tiles.StateOK:
		g.Money -= 850
	}
}
