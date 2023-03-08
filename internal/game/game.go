package game

import (
	"fmt"

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
	Engines    *Engines

	screenWidth  int
	screenHeight int
}

func NewGame() *Game {
	game := &Game{
		Input:      NewInput(),
		Background: NewBackground(),
		Engines:    NewEngines(),
	}

	return game
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.Draw(screen)

	// do an FPS count debug print on the top left corner
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", ebiten.ActualFPS()), 0, 0)
	// NOTE: use text.Draw(...) to print normal text (like a game menu or whatever)

	g.Engines.Draw(screen)
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	g.screenWidth = outsideWidth
	g.screenHeight = outsideHeight

	return g.screenWidth, g.screenHeight
}

// Update implements ebiten.Game
func (g *Game) Update() (err error) {
	err = g.Engines.Update(g.Input)

	return
}
