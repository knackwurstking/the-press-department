package game

import "github.com/hajimehoshi/ebiten/v2"

type GameComponentConfig interface {
	InputConfig | BackgroundConfig | EnginesConfig
}

type GameComponent[T GameComponentConfig] interface {
	SetGame(game *Game)
	SetConfig(config *T)
	GetConfig() *T
	Draw(screen *ebiten.Image)
	Update() error
}
