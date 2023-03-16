package game

import "github.com/hajimehoshi/ebiten/v2"

type GameComponentConfig interface {
	InputConfig | BackgroundConfig
}

type GameComponent[T GameComponentConfig] interface {
	SetGame(game *Game)
	SetConfig(config *T)
	GetConfig() *T
	Update() error
	Draw(screen *ebiten.Image)
}
