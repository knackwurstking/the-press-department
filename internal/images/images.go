package images

import (
	_ "embed"
)

var (
	//go:embed tile-0.png
	Tile []byte

	//go:embed tile-1.png
	TileWithCrack []byte
)
