package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(940, 470)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("The Press Department")

	if err := ebiten.RunGame(NewGame(DefaultScale * 1.5)); err != nil {
		log.Fatalf("Run game failed: %s", err.Error())
	}
}
