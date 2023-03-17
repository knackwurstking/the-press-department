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
