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

	_debugCounter string
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

	g.debugFPS(screen)
	g.debugEngines(screen)

	g.Engines.Draw(screen)
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	g.ScreenWidth = outsideWidth
	g.ScreenHeight = outsideHeight

	return g.ScreenWidth, g.ScreenHeight
}

// Update implements ebiten.Game
func (g *Game) Update() error {
	return g.Engines.Update(g.Input)
}

func (g *Game) debugEngines(screen *ebiten.Image) {
	g._debugCounter = fmt.Sprintf(
		"Press Speed: %.1fh | Tiles Produced: %06d | RB: %2d",
		g.Engines.BPM,
		g.Engines.tilesCount,
		len(g.Engines.tiles),
	)

	ebitenutil.DebugPrintAt(
		screen,
		g._debugCounter,
		g.ScreenWidth-len(g._debugCounter)*6,
		0,
	)
}

func (g *Game) debugFPS(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", ebiten.ActualFPS()), 0, 0)
}
