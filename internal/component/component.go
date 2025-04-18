package component

import "github.com/hajimehoshi/ebiten/v2"

type ComponentData interface {
	BackgroundData | RollerConveyorData | RollerConveyorUserInputData | ToolData
}

type Component[T ComponentData] interface {
	Data() *T
	Layout(outsideWidth, outsideHeight int) (int, int)
	Draw(screen *ebiten.Image)
	Update() error
}

type Coord struct {
	X, Y float64
}
