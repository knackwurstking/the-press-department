package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{}

// Layout implements ebiten.Game
func (*Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Update implements ebiten.Game
func (*Game) Update(screen *ebiten.Image) (err error) {
	err = ebitenutil.DebugPrint(screen, "The Press Department")

	return
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("The Press Department")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalf("Run game failed: %s", err.Error())
	}
}
