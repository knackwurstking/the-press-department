package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game ...
//
// Input 			- handles the controls (with mobile touch support)
// Background - handles the game background
// TODO: where to put the engines and the tiles (need some root struct like "Foreground" or "?")
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
	ebiten.SetWindowSize(940, 470)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("The Press Department")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalf("Run game failed: %s", err.Error())
	}
}
