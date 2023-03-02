package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"the-press-department/internal/game"
)

func main() {
	ebiten.SetWindowSize(940, 470)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("The Press Department")

	g := game.NewGame(
		game.NewInput(),
		game.NewBackground(),
		game.NewPress(),
		game.NewEngines(),
	)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("Run game failed: %s", err.Error())
	}
}
