package game

import "github.com/hajimehoshi/ebiten/v2"

type GameComponentConfig interface {
	InputConfig | BackgroundConfig | EnginesConfig | ConveyorConfig
}

type GameComponent[T GameComponentConfig] interface {
	SetConfig(config *T)
	GetConfig() *T
	Layout(outsideWidth, outsideHeight int) (int, int)
	Draw(screen *ebiten.Image)
	Update() error
}
