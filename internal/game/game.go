package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game ...
//
// Input 			- handles the controls (with mobile touch support)
// Background - handles the game background
// Press 			- produces tiles and outputs each tile to the Engines
// Engines 		- transports the tiles (from the Press) from A to B
type Game struct {
	Input      *Input
	Background *Background
	Press      *Press
	Engines    *Engines
}

func NewGame(i *Input, b *Background, p *Press, e *Engines) *Game {
	game := &Game{
		Input:      i,
		Background: b,
		Press:      p,
		Engines:    e,
	}

	return game
}

// Layout implements ebiten.Game
func (*Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Update implements ebiten.Game
func (*Game) Update(screen *ebiten.Image) (err error) {
	err = ebitenutil.DebugPrint(screen, "The Press Department")

	return
}
