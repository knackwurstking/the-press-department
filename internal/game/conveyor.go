package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ConveyorConfig struct {
	Sprite     *RollSprite
	Scale      *float64
	Hz         float64
	HzMultiply float64

	X, Y float64

	// Update fields
	rSum, r, x, y, size float64
}

func (c *ConveyorConfig) SetUpdateData(r, x, y, size float64) {
	c.r = r
	c.x = x
	c.y = y
	c.size = size
}

func (c *ConveyorConfig) SetSprite() {
	c.rSum += c.r
	w, _ := c.Sprite.GetAssetSize()
	if c.rSum >= w {
		c.Sprite.NextSprite()
		c.rSum = 0
	}
}

func (c *ConveyorConfig) GetHeight() float64 {
	_, h := c.Sprite.GetAssetSize()
	return h
}

type Conveyor struct {
	game   *Game
	config *ConveyorConfig
	rolls  []Coord
}

func NewConveyor(config *ConveyorConfig) *Conveyor {
	c := &Conveyor{
		config: config,
		rolls:  make([]Coord, 0),
	}

	return c
}

func (c *Conveyor) Draw(screen *ebiten.Image) {
	for i := 0; i < len(c.rolls); i++ {
		c.config.Sprite.Draw(screen, c.rolls[i].X, c.rolls[i].Y)
	}
}

func (c *Conveyor) Update() error {
	c.config.X = c.config.x
	c.config.Y = c.config.y

	c.config.SetSprite()
	w, _ := c.config.Sprite.GetAssetSize()
	padding := w * 3

	c.rolls = make([]Coord, 0)
	for p := c.config.X; p <= c.config.size; p += (w + padding) {
		c.rolls = append(c.rolls, Coord{X: float64(p), Y: c.config.Y})
	}

	return nil
}

func (c *Conveyor) SetGame(game *Game) {
	c.game = game
}

func (c *Conveyor) SetConfig(config *ConveyorConfig) {
	c.config = config
}

func (c *Conveyor) GetConfig() *ConveyorConfig {
	return c.config
}
