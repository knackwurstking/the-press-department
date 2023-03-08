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
	Input        *Input
	Background   *Background
	Engines      *Engines
	ScreenWidth  int
	ScreenHeight int
}

func NewGame() *Game {
	game := &Game{
		Input:      NewInput(),
		Background: NewBackground(),
		Engines:    NewEngines(),
	}

	// pass game pointer to the engine
	game.Engines.Game = game

	return game
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.Draw(screen)

	// do an FPS count debug print on the top left corner
	// NOTE: use text.Draw(...) to print normal text (like a game menu or whatever)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", ebiten.ActualFPS()), 0, 0)

	s := fmt.Sprintf("Tiles Count: %5d; RB: %2d", g.Engines.tilesCount, len(g.Engines.tiles))
	ebitenutil.DebugPrintAt(
		screen,
		s,
		g.ScreenWidth-len(s)*6,
		0,
	)

	g.Engines.Draw(screen)
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	g.ScreenWidth = outsideWidth
	g.ScreenHeight = outsideHeight

	return g.ScreenWidth, g.ScreenHeight
}

// Update implements ebiten.Game
func (g *Game) Update() (err error) {
	err = g.Engines.Update(g.Input)

	return
}
