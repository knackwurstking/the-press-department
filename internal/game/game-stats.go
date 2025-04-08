package game

import "the-press-department/internal/tiles"

// Stats for saving the game state
type GameStats struct {
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

	// ConveyorHz (setup value)
	ConveyorHz float64 `json:"conveyor-hz"`

	// ConveyorHzMultiply (setup value)
	ConveyorHzMultiply float64 `json:"conveyor-hz-multiply"`
}

func (s *GameStats) AddGoodTile() {
	s.GoodTiles++
	s.Money += 150
}

func (s *GameStats) AddBadTile() {
	s.BadTiles++
	s.Money -= 450
}

// "add thrown away good tile", "add thrown away bad tile"
func (s *GameStats) AddThrownAwayTile(tile tiles.Tiles) {
	if !tile.IsThrownAway() {
		return
	}

	// check if tile was ok (if ok then add a penalty "-1000$")...
	switch tile.Data().State {
	case tiles.StateCrack:
		s.Money -= 50
	case tiles.StateOK:
		s.Money -= 850
	}
}
