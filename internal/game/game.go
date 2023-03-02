package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game controls all the game logic
//
// Input 			- handles the controls (with mobile touch support)
// Background - handles the game background
// Press 			- produces tiles and outputs each tile to the Engines
// Engines 		- transports the tiles (from the Press) from A to B
type Game struct {
	Input      *Input
	Background *Background
	Board      *Board

	screenWidth  int
	screenHeight int
}

func NewGame() *Game {
	game := &Game{
		Input:      &Input{},
		Background: &Background{},
		Board:      NewBoard(),
	}

	return game
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.Draw(screen)

	ebitenutil.DebugPrint(screen, "The Press Department")
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	g.screenWidth = outsideWidth
	g.screenHeight = outsideHeight

	return g.screenWidth, g.screenHeight
}

// Update implements ebiten.Game
func (g *Game) Update() error {
	err := g.Board.Update(g.Input)

	return err
}
