package stats

import "the-press-department/internal/tiles"

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

	// RollingRailwayHz (setup value)
	RollingRailwayHz float64 `json:"conveyor-hz"`

	// RollingRailwayHzMultiply (setup value)
	RollingRailwayHzMultiply float64 `json:"conveyor-hz-multiply"`
}

func (g *Game) AddGoodTile() {
	g.GoodTiles++
	g.Money += 150
}

func (g *Game) AddBadTile() {
	g.BadTiles++
	g.Money -= 450
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
	case tiles.StateOK:
		g.Money -= 850
	}
}
