package game

import "github.com/hajimehoshi/ebiten/v2"

type ComponentData interface {
	BackgroundData | EnginesData | EnginesInputData | ConveyorData
}

type Component[T ComponentData] interface {
	SetData(config *T)
	GetData() *T
	Layout(outsideWidth, outsideHeight int) (int, int)
	Draw(screen *ebiten.Image)
	Update() error
}
