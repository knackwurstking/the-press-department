package images

import (
	_ "embed"
)

var (
	//go:embed ground-0.png
	Ground []byte

	//go:embed tile-0.png
	Tile []byte

	//go:embed tile-1.png
	TileWithCrack []byte

	//go:embed roll-0.1.png
	Roll0 []byte

	//go:embed roll-0.2.png
	Roll1 []byte

	//go:embed roll-0.3.png
	Roll2 []byte
)
