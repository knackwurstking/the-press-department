package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	DefaultScale float64 = 0.1

	// modes
	ModePause = Mode(1)
	ModeGame  = Mode(2)

	FontDPI = float64(71)

	FontSize = float64(24)
	FontFace font.Face

	FontSizeSmall = float64(16)
	FontFaceSmall font.Face

	FontSizeBig = float64(31)
	FontFaceBig font.Face
)

func init() {
	ttf, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		panic(err)
	}

	FontFace, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    FontSize,
		DPI:     FontDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	FontFaceSmall, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    FontSizeSmall,
		DPI:     FontDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	FontFaceBig, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    FontSizeBig,
		DPI:     FontDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
}

type Mode int

// Stats for saving the game state
type Stats struct {
	// Money holds the number of your money
	Money int `json:"money"`

	// GoodTiles shows the number of good tiles passed
	GoodTiles int `json:"good-tiles"`

	// BadTiles shows the number of bad tiles passed
	BadTiles int `json:"bad-tiles"`

	// TilesProduced is the number of tiles produced from press
	TilesProduced int `json:"tiles-produced"`

	// PressBPM (setup value)
	PressBPM float64 `json:"press-BPM"`

	// ConveyorHz (setup value)
	ConveyorHz float64 `json:"conveyor-hz"`

	// ConveyorHzMultiply (setup value)
	ConveyorHzMultiply float64 `json:"conveyor-hz-multiply"`
}

// Game controls all the game logic
type Game struct {
	Mode       Mode
	Background GameComponent[BackgroundConfig]
	Engines    GameComponent[EnginesConfig]

	Stats Stats

	screenWidth, screenHeight int
	scale                     float64

	lastUpdate time.Time
}

func NewGame(scale float64) *Game {
	game := &Game{
		Mode: ModePause,
		Stats: Stats{
			TilesProduced:      0, // Engines tilesProduced config field
			PressBPM:           6.5,
			ConveyorHz:         8.0,
			ConveyorHzMultiply: 2.5,
		},
		Background: NewBackground(&BackgroundConfig{
			Scale: scale,
			Image: ebiten.NewImageFromImage(ImageGround),
		}),
		Engines: NewEngines(&EnginesConfig{
			Input: NewEnginesInput(&EnginesInputConfig{}),
			Scale: scale,
		}),
		scale: scale,
	}

	game.Engines.GetConfig().SetBPM(&game.Stats.PressBPM)
	game.Engines.GetConfig().SetHz(&game.Stats.ConveyorHz)
	game.Engines.GetConfig().SetHzMultiply(&game.Stats.ConveyorHzMultiply)
	game.Engines.GetConfig().SetTilesCount(&game.Stats.TilesProduced)

	return game
}

// Layout implements ebiten.Game
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	g.screenWidth = outsideWidth
	g.screenHeight = outsideHeight

	g.Background.Layout(outsideWidth, outsideHeight)
	g.Engines.Layout(outsideWidth, outsideHeight)

	return outsideWidth, outsideHeight
}

// Update implements ebiten.Game
func (g *Game) Update() error {
	// catch standby and pause the game (check for standby via time package)
	if time.Since(g.lastUpdate) >= time.Second {
		g.Mode = ModePause
	}

	g.Background.GetConfig().Scale = g.scale
	g.Engines.GetConfig().Scale = g.scale

	switch g.Mode {
	case ModePause:
		g.Engines.GetConfig().Pause = true
		// Listen for keys to continue (or start the game)
		if g.isKeyPressed() {
			// Continue or start the game
			g.Engines.GetConfig().Pause = false
			g.Mode = ModeGame
		}
	case ModeGame:
	}

	_ = g.Background.Update()
	_ = g.Engines.Update()

	g.updateStats()
	g.lastUpdate = time.Now()

	return nil
}

// Draw implements ebiten.Game
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Mode {
	case ModePause:
		g.drawPause(screen)
	case ModeGame:
		g.drawGame(screen)
	}
}

func (g *Game) updateStats() {
	// TODO: update game status (put everything in g.Stats)
}

func (g *Game) isKeyPressed() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}

	touchIDs := make([]ebiten.TouchID, 0)
	touchIDs = inpututil.AppendJustPressedTouchIDs(touchIDs[:0])
	return len(touchIDs) > 0
}

func (g *Game) drawPause(screen *ebiten.Image) {
	g.Background.Draw(screen)
	g.Engines.Draw(screen)

	titleTexts := []string{
		"PAUSE",
	}

	texts := []string{
		"Click (or Touch) to start.",
	}

	for i, l := range titleTexts {
		x := int((g.screenWidth - len(l)*int(FontSizeBig)) / 2)
		y := (i + 4) * int(FontSizeBig)
		text.Draw(screen, l, FontFaceBig, x, y, color.White)
	}

	for i, l := range texts {
		x := int((g.screenWidth - len(l)*int(FontSizeSmall)) / 2)
		y := ((len(titleTexts) + 3) * int(FontSizeBig)) + ((i + 4) * int(FontSizeSmall))
		text.Draw(screen, l, FontFaceSmall, x, y, color.White)
	}

	g.drawDebug(screen)
}

func (g *Game) drawGame(screen *ebiten.Image) {
	// run the game
	g.Background.Draw(screen)
	g.Engines.Draw(screen)
	g.drawStats(screen)
	g.drawDebug(screen)
}

func (g *Game) drawStats(screen *ebiten.Image) {
	// TODO: Drawing stats like +/- money for each tile... (Need a State object with json support for saving)
	//  - bottom/left corner: money in the bank
	//  - +<dollar>, green
	//  - -<dollar>, red
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	// debug overlay: "FPS"
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", ebiten.ActualFPS()), 0, 0)

	// debug overlay: "Engines Info"
	// 1. Row
	counter := fmt.Sprintf("Press Speed: %.1fh", g.Engines.GetConfig().GetBPM())
	ebitenutil.DebugPrintAt(screen, counter, g.screenWidth-(len(counter)*6+2), 0)

	// 2. Row
	counter = fmt.Sprintf("Tiles Produced: %d",
		g.Engines.GetConfig().GetTilesCount())
	ebitenutil.DebugPrintAt(screen, counter, g.screenWidth-(len(counter)*6+2), 16)

	// 3. Row
	counter = fmt.Sprintf("RB: %d [%.1f hz]", len(g.Engines.GetConfig().GetTiles()),
		g.Engines.GetConfig().GetHz())
	ebitenutil.DebugPrintAt(screen, counter, g.screenWidth-(len(counter)*6+2), 32)
}
