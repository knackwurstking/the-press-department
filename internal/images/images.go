package images

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed ground-0.png
	ground []byte
	Ground *ebiten.Image

	//go:embed tile-0.png
	tileStateOK []byte
	TileStateOK *ebiten.Image

	//go:embed tile-1.png
	tileStateCrack []byte
	TileStateCrack *ebiten.Image

	//go:embed tile-2.png
	tileStateStampAdhesive []byte
	TileStateStampAdhesive *ebiten.Image

	//go:embed roll-0_0.png
	roll0 []byte
	Roll0 *ebiten.Image

	//go:embed roll-0_1.png
	roll1 []byte
	Roll1 *ebiten.Image

	//go:embed roll-0_2.png
	roll2 []byte
	Roll2 *ebiten.Image

	//go:embed roll-0_3.png
	roll3 []byte
	Roll3 *ebiten.Image

	//go:embed roll-0_4.png
	roll4 []byte
	Roll4 *ebiten.Image

	//go:embed roll-0_5.png
	roll5 []byte
	Roll5 *ebiten.Image

	//go:embed roll-0_6.png
	roll6 []byte
	Roll6 *ebiten.Image

	//go:embed roll-0_7.png
	roll7 []byte
	Roll7 *ebiten.Image
)

func init() {
	// Ground
	i, _, err := image.Decode(bytes.NewReader(ground))
	if err != nil {
		panic(err)
	}
	Ground = ebiten.NewImageFromImage(i)

	// TileOK
	f, _, err := image.Decode(bytes.NewReader(tileStateOK))
	if err != nil {
		panic(err)
	}
	TileStateOK = ebiten.NewImageFromImage(f)

	// TileCrack
	f, _, err = image.Decode(bytes.NewReader(tileStateCrack))
	if err != nil {
		panic(err)
	}
	TileStateCrack = ebiten.NewImageFromImage(f)

	// Stamp Adhesive
	f, _, err = image.Decode(bytes.NewReader(tileStateStampAdhesive))
	if err != nil {
		panic(err)
	}
	TileStateStampAdhesive = ebiten.NewImageFromImage(f)

	// Roll 0-7
	img, _, err := image.Decode(bytes.NewReader(roll0))
	if err != nil {
		panic(err)
	}
	Roll0 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(roll1))
	if err != nil {
		panic(err)
	}
	Roll1 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(roll2))
	if err != nil {
		panic(err)
	}
	Roll2 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(roll3))
	if err != nil {
		panic(err)
	}
	Roll3 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(roll4))
	if err != nil {
		panic(err)
	}
	Roll4 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(roll5))
	if err != nil {
		panic(err)
	}
	Roll5 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(roll6))
	if err != nil {
		panic(err)
	}
	Roll6 = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(roll7))
	if err != nil {
		panic(err)
	}
	Roll7 = ebiten.NewImageFromImage(img)
}
