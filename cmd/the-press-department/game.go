//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"image/color"
	"the-press-department/internal/component"
	"the-press-department/internal/stats"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	DefaultScale float64 = 0.1

	// modes
	ModePause   = Mode(1)
	ModeGame    = Mode(2)
	ModeSuspend = Mode(3)

	FontDPI = float64(71)

	FontSize = float64(24)
	FontFace text.Face

	FontSizeSmall = float64(16)
	FontFaceSmall text.Face

	FontSizeBig = float64(31)
	FontFaceBig text.Face
)

func init() {
	ttf, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		panic(err)
	}

	f, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    FontSize,
		DPI:     FontDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
	FontFace = text.NewGoXFace(f)

	f, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    FontSizeSmall,
		DPI:     FontDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
	FontFaceSmall = text.NewGoXFace(f)

	f, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    FontSizeBig,
		DPI:     FontDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
	FontFaceBig = text.NewGoXFace(f)
}

type Mode int

// Game controls all the game logic
type Game struct {
	Mode           Mode
	Background     component.Component[component.BackgroundData]
	RollingRailway component.Component[component.RollingRailwayData]

	Stats *stats.Game

	screenWidth, screenHeight int
	scale                     float64

	lastUpdate time.Time
}

func NewGame(scale float64) *Game {
	stats := &stats.Game{
		TilesProduced:            0, // Engine tilesProduced config field
		PressBPM:                 6.5,
		RollingRailwayHz:         8.0,
		RollingRailwayHzMultiply: 2.5,
	}

	game := &Game{
		Mode:  ModePause,
		Stats: stats,
		Background: component.NewBackground(&component.BackgroundData{
			Scale: scale,
			Image: ebiten.NewImageFromImage(component.ImageGround),
		}),
		RollingRailway: component.NewRollingRailway(&component.RollingRailwayData{
			Scale:  scale,
			Sprite: component.NewRollSprite(&scale),
		}),
		scale: scale,
	}

	return game
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	g.screenWidth = outsideWidth
	g.screenHeight = outsideHeight

	g.Background.Layout(outsideWidth, outsideHeight)
	g.RollingRailway.Layout(outsideWidth, outsideHeight)

	return outsideWidth, outsideHeight
}

// Update implements ebiten.Game
func (g *Game) Update() error {
	// NOTE: Disable ModeSuspend for now
	//
	//if time.Since(g.lastUpdate) >= time.Second && g.Mode != ModePause {
	//	g.Mode = ModeSuspend
	//}

	g.Background.Data().Scale = g.scale
	g.RollingRailway.Data().Scale = g.scale

	switch g.Mode {
	case ModePause, ModeSuspend:
		g.RollingRailway.Data().Pause = true
		// Listen for keys to continue (or start the game)
		if g.isKeyPressed() {
			// Continue or start the game
			g.RollingRailway.Data().Pause = false
			g.Mode = ModeGame
		}
	case ModeGame:
	}

	_ = g.Background.Update()
	_ = g.RollingRailway.Update()

	g.lastUpdate = time.Now()

	return nil
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Mode {
	case ModePause, ModeSuspend:
		g.drawPause(screen)
	case ModeSuspend:
	case ModeGame:
		g.drawGame(screen)
	}
}

func (g *Game) isKeyPressed() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}

	var touchIDs []ebiten.TouchID
	touchIDs = inpututil.AppendJustPressedTouchIDs(touchIDs)
	return len(touchIDs) > 0
}

func (g *Game) drawPause(screen *ebiten.Image) {
	g.Background.Draw(screen)
	g.RollingRailway.Draw(screen)

	g.drawStats(screen)

	titleTexts := []string{
		"PAUSE",
	}

	var texts []string
	switch g.Mode {
	case ModePause:
		texts = append(texts, "Click to start.")
	case ModeSuspend:
		texts = append(texts, "Click to continue.")
	}

	for i, l := range titleTexts {
		x := (float64(g.screenWidth) - float64(len(l))*FontSizeBig) / 2
		y := float64(i+4) * FontSizeBig

		// text.Draw(screen, l, FontFaceBig, x, y, color.White)
		dopt := &text.DrawOptions{}
		dopt.ColorScale.ScaleWithColor(color.White)
		dopt.GeoM.Translate(x, y)
		text.Draw(screen, l, FontFaceBig, dopt)
	}

	for i, l := range texts {
		x := (float64(g.screenWidth) - float64(len(l))*FontSizeSmall) / 2
		y := (float64(len(titleTexts)+3) * FontSizeBig) + (float64(i+4) * FontSizeSmall)

		dopt := &text.DrawOptions{}
		dopt.ColorScale.ScaleWithColor(color.White)
		dopt.GeoM.Translate(x, y)
		text.Draw(screen, l, FontFaceSmall, dopt)
	}

	g.drawDebug(screen)
}

func (g *Game) drawGame(screen *ebiten.Image) {
	// run the game
	g.Background.Draw(screen)
	g.RollingRailway.Draw(screen)

	g.drawStats(screen)
	g.drawDebug(screen)
}

func (g *Game) drawStats(screen *ebiten.Image) {
	// Draw "$: <n>"
	textMoney := fmt.Sprintf("%d$", g.Stats.Money)
	c := color.RGBA{255, 0, 0, 255}
	if g.Stats.Money >= 0 {
		c = color.RGBA{0, 255, 0, 255}
	}

	dopt := &text.DrawOptions{}
	dopt.ColorScale.ScaleWithColor(c)
	dopt.GeoM.Translate(1, float64(g.screenHeight))
	text.Draw(screen, textMoney, FontFace, dopt)
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	// debug overlay: "FPS"
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", ebiten.ActualFPS()), 0, 0)

	// debug overlay: "Engine Info"
	// 1. Row
	counter := fmt.Sprintf("Press Speed: %.1fh", g.RollingRailway.Data().PressBPM())
	ebitenutil.DebugPrintAt(screen, counter, g.screenWidth-(len(counter)*6+2), 0)

	// 2. Row
	counter = fmt.Sprintf("Tiles Produced: %d", g.Stats.TilesProduced)
	ebitenutil.DebugPrintAt(screen, counter, g.screenWidth-(len(counter)*6+2), 16)

	// 3. Row
	counter = fmt.Sprintf("RB: %d [%.1f hz]", len(g.RollingRailway.Data().GetTiles()),
		g.RollingRailway.Data().GetHz())
	ebitenutil.DebugPrintAt(screen, counter, g.screenWidth-(len(counter)*6+2), 32)
}
