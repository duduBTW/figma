package layer

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	x_KEY = "x"
	y_KEY = "y"
	empty = "_$_$_$_$_empty__._$$__"
)

type AnimatedVector2 struct {
	Id string
	X  AnimatedProp
	Y  AnimatedProp
}

func NewAnimatedVector2(id string, x float32, y float32) AnimatedVector2 {
	return AnimatedVector2{
		Id: id,
		X:  NewAnimatedProp(x, x_KEY),
		Y:  NewAnimatedProp(y, y_KEY),
	}
}

func (v2 *AnimatedVector2) Controls(ui *lib.UIStruct, comp components.Components) lib.Component {
	return func(avaliablePosition rl.Rectangle) (func(), float32, float32) {
		row := NewControlsLayout(avaliablePosition)
		row.Add(Label("Position"))
		row.Add(v2.Inputs(ui, comp))
		return row.Draw, 0, row.Size.Height
	}
}

func (v2 *AnimatedVector2) Inputs(ui *lib.UIStruct, comp components.Components) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := InputsLayout(2, rect)
		row.Add(v2.X.Input(ui, comp))
		row.Add(v2.Y.Input(ui, comp))
		return row.Draw, 0, row.Size.Height
	}
}

func (v2 *AnimatedVector2) CanDrawTimeline() bool {
	return v2.X.CanDrawTimeline() || v2.Y.CanDrawTimeline()
}

func (v2 *AnimatedVector2) Timeline(layout *lib.Layout, ui *lib.UIStruct, comp components.Components) {
	if v2.X.CanDrawTimeline() {
		layout.Add(comp.TimelineRow("Position x", v2.X.NewInput(ui, comp), v2.X.SortedKeyframes))
	}

	if v2.Y.CanDrawTimeline() {
		layout.Add(comp.TimelineRow("Position y", v2.Y.NewInput(ui, comp), v2.Y.SortedKeyframes))
	}
}
