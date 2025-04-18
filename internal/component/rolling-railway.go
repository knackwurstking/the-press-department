package component

import (
	"the-press-department/internal/stats"
	"the-press-department/internal/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type RollingRailway struct {
	Engine Component[EngineData]

	stats                     *stats.Game
	data                      *RollingRailwayData
	rolls                     []Coord
	screenWidth, screenHeight float64
}

func NewRollingRailway(data *RollingRailwayData) Component[RollingRailwayData] {
	c := &RollingRailway{
		Engine: NewEngine(&EngineData{
			Stats: data.Stats,
			Scale: &data.Scale,
		}),
		data:  data,
		rolls: make([]Coord, 0),
	}

	return c
}

func (c *RollingRailway) Layout(outsideWidth, outsideHeight int) (int, int) {
	c.screenWidth = float64(outsideWidth)
	c.screenHeight = float64(outsideHeight)

	c.Engine.Layout(outsideWidth, outsideHeight)

	return outsideWidth, outsideHeight
}

func (c *RollingRailway) Update() error {
	c.data.X = c.data.x
	c.data.Y = c.data.y

	c.data.SetSprite()
	w, _ := c.data.Sprite.GetAssetSize()
	padding := w * 3

	c.rolls = make([]Coord, 0)
	for p := c.data.X; p <= c.data.size; p += (w + padding) {
		c.rolls = append(c.rolls, Coord{X: float64(p), Y: c.data.Y})
	}

	return c.Engine.Update()
}

func (c *RollingRailway) Draw(screen *ebiten.Image) {
	for i := 0; i < len(c.rolls); i++ {
		c.data.Sprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}
	c.Engine.Draw(screen)
}

func (c *RollingRailway) Data() *RollingRailwayData {
	return c.data
}

type RollingRailwayData struct {
	Stats  *stats.Game
	Sprite *RollSprite
	Scale  float64
	Pause  bool // Pause will stop the machines :)

	X, Y float64

	// Update fields
	rSum, r, x, y, size float64

	tiles []tiles.Tiles
}

func (c *RollingRailwayData) SetUpdateData(r, x, y, size float64) {
	c.r = r
	c.x = x
	c.y = y
	c.size = size
}

func (c *RollingRailwayData) SetSprite() {
	c.rSum += c.r
	w, _ := c.Sprite.GetAssetSize()
	if c.rSum >= w {
		c.Sprite.NextSprite()
		c.rSum = 0
	}
}

func (c *RollingRailwayData) GetHeight() float64 {
	_, h := c.Sprite.GetAssetSize()
	return h
}

func (c *RollingRailwayData) PressBPM() float64 {
	if c.Pause {
		return 0
	}

	return c.Stats.PressBPM
}

func (c *RollingRailwayData) GetHz() float64 {
	if c.Pause {
		return 0
	}

	return c.Stats.RollingRailwayHz
}

func (c *RollingRailwayData) GetTiles() []tiles.Tiles {
	return c.tiles
}
