package game

// Stats for saving the game state
type Stats struct {
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

func (s *Stats) AddGoodTile() {
	s.GoodTiles++
	s.Money += 100
}

func (s *Stats) AddBadTile() {
	s.BadTiles++
	s.Money -= 250
}

// TODO: "add thrown away good tile", "add thrown away bad tile"
func (s *Stats) AddThrownAwayTile(tile *Tile) {
	if !tile.IsThrownAway() {
		return
	}

	// TODO: check if tile was ok (if ok then add a penalty "-1000$")...
	//switch tile.State {
	//case StateCrack:
	//	s.Money -= 50
	//case StateOK:
	//	s.Money -= 1000
	//}
}
