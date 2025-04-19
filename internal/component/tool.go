package component

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ToolTypeMain = "main"
)

type ToolType string

type ToolData struct {
	t ToolType
}

func (td *ToolData) Type() ToolType {
	return td.t
}

type Tool struct {
	ToolData

	width  float64
	height float64
}

func NewTool(t ToolType) Component[ToolData] {
	return &Tool{
		ToolData: ToolData{
			t: t,
		},
	}
}

func (t Tool) Layout(width, height int) (w int, h int) {
	t.width = float64(width)
	t.height = float64(height)

	return width, height
}

func (t Tool) Update() error {
	// ...

	return fmt.Errorf("under construction")
}

func (t Tool) Draw(i *ebiten.Image) {
	// ...
}

func (t Tool) Data() *ToolData {
	return &t.ToolData
}
