package game

// Board holds all the data and coordinates (like tiles positions and engine positions)
type Board struct {
	Press   *Press
	Engines *Engines
}

func NewBoard() *Board {
	return &Board{
		Press:   &Press{},
		Engines: &Engines{},
	}
}

func (*Board) Update(input *Input) error {
	return nil
}
