package game

type Roll struct {
	Scale *float64
}

func NewRoll(scale *float64) Roll {
	return Roll{
		Scale: scale,
	}
}
