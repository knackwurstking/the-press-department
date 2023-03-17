package game

import "github.com/hajimehoshi/ebiten/v2"

type GameComponentConfig interface {
	BackgroundConfig | EnginesConfig | EnginesInputConfig | ConveyorConfig
}

type GameComponent[T GameComponentConfig] interface {
	SetConfig(config *T)
	GetConfig() *T
	Layout(outsideWidth, outsideHeight int) (int, int)
	Draw(screen *ebiten.Image)
	Update() error
}
