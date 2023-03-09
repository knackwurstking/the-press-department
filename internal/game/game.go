package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	DefaultScale float64 = 0.15
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

	scale float64

	_debugCounter string
}

func NewGame(scale float64) *Game {
	game := &Game{
		Input:      NewInput(),
		Background: NewBackground(scale, ImageGround),
		Engines:    NewEngines(),
		scale:      scale,
	}

	// pass game pointer to the engine
	game.Engines.Game = game

	return game
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.Draw(screen, float64(g.ScreenWidth), float64(g.ScreenHeight))
	g.Engines.Draw(screen)

	g.debugFPS(screen)
	g.debugEngines(screen)
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	g.ScreenWidth = outsideWidth
	g.ScreenHeight = outsideHeight

	return g.ScreenWidth, g.ScreenHeight
}

// Update implements ebiten.Game
func (g *Game) Update() error {
	if g.Engines.Scale != g.scale {
		g.Engines.Scale = g.scale
		g.Background.Scale = g.scale
	}

	return g.Engines.Update(g.Input)
}

func (g *Game) GetScale() float64 {
	return g.scale
}

func (g *Game) SetScale(f float64) {
	g.scale = f
}

func (g *Game) debugEngines(screen *ebiten.Image) {
	// 1. Row
	g._debugCounter = fmt.Sprintf(
		"Press Speed: %.1fh",
		g.Engines.BPM,
	)

	ebitenutil.DebugPrintAt(
		screen,
		g._debugCounter,
		g.ScreenWidth-(len(g._debugCounter)*6+2),
		0,
	)

	// 2. Row
	g._debugCounter = fmt.Sprintf(
		"Tiles Produced: %d",
		g.Engines.tilesCount,
	)

	ebitenutil.DebugPrintAt(
		screen,
		g._debugCounter,
		g.ScreenWidth-(len(g._debugCounter)*6+2),
		16,
	)

	// 3. Row
	g._debugCounter = fmt.Sprintf(
		"RB: %d [%.1f hz]",
		len(g.Engines.tiles),
		g.Engines.Hz,
	)

	ebitenutil.DebugPrintAt(
		screen,
		g._debugCounter,
		g.ScreenWidth-(len(g._debugCounter)*6+2),
		32,
	)
}

func (g *Game) debugFPS(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", ebiten.ActualFPS()), 0, 0)
}
