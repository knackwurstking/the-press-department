package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	DefaultScale float64 = 0.1
)

// Game controls all the game logic
//
// Input 			- handles the controls (with mobile touch support)
// Background - handles the game background
// Press 			- produces tiles and outputs each tile to the Engines
// Engines 		- transports the tiles (from the Press) from A to B
type Game struct {
	Input      GameComponent[InputConfig]
	Background GameComponent[BackgroundConfig]
	Engines    GameComponent[EnginesConfig]

	screenWidth  int
	screenHeight int
	scale        float64
}

func NewGame(scale float64) *Game {
	game := &Game{
		Input: NewInput(&InputConfig{}),
		Background: NewBackground(&BackgroundConfig{
			Scale: scale,
			Image: ebiten.NewImageFromImage(ImageGround),
		}),
		Engines: NewEngines(&EnginesConfig{
			Scale:      scale,
			BPM:        6.5,
			Hz:         8,
			HzMultiply: 2.5,
		}),
		scale: scale,
	}

	// pass game pointer to the engine
	game.Engines.GetConfig().Input = game.Input

	return game
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	g.screenWidth = outsideWidth
	g.screenHeight = outsideHeight

	g.Input.Layout(g.screenWidth, g.screenHeight)
	g.Background.Layout(g.screenWidth, g.screenHeight)
	g.Engines.Layout(g.screenWidth, g.screenHeight)

	return g.screenWidth, g.screenHeight
}

// Update implements ebiten.Game
func (g *Game) Update() error {
	g.Engines.GetConfig().Scale = g.scale
	g.Background.GetConfig().Scale = g.scale

	_ = g.Input.Update()
	_ = g.Background.Update()
	return g.Engines.Update()
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	g.Background.Draw(screen)
	g.Engines.Draw(screen)

	g.debugFPS(screen)
	g.debugEngines(screen)
}

func (g *Game) GetScale() float64 {
	return g.scale
}

func (g *Game) SetScale(f float64) {
	g.scale = f
}

func (g *Game) debugEngines(screen *ebiten.Image) {
	// 1. Row
	counter := fmt.Sprintf(
		"Press Speed: %.1fh",
		g.Engines.GetConfig().BPM,
	)

	ebitenutil.DebugPrintAt(
		screen,
		counter,
		g.screenWidth-(len(counter)*6+2),
		0,
	)

	// 2. Row
	counter = fmt.Sprintf(
		"Tiles Produced: %d",
		g.Engines.GetConfig().GetTilesCount(),
	)

	ebitenutil.DebugPrintAt(
		screen,
		counter,
		g.screenWidth-(len(counter)*6+2),
		16,
	)

	// 3. Row
	counter = fmt.Sprintf(
		"RB: %d [%.1f hz]",
		len(g.Engines.GetConfig().GetTiles()),
		g.Engines.GetConfig().Hz,
	)

	ebitenutil.DebugPrintAt(
		screen,
		counter,
		g.screenWidth-(len(counter)*6+2),
		32,
	)
}

func (g *Game) debugFPS(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", ebiten.ActualFPS()), 0, 0)
}
