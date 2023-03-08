package game

type Tile struct {
	Crack    bool          // Crack holds whenever this tile has a crack or not
	Position *TilePosition // X holds the position where the tile is moving on the engine
}

func NewTile(position *TilePosition) *Tile {
	return &Tile{
		Crack:    false,
		Position: position, // tile not exists on engine
	}
}

type TilePosition struct {
	X int64
}

func NewTilePosition(x int64) *TilePosition {
	return &TilePosition{
		X: x,
	}
}
