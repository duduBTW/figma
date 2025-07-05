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

func (v2 *AnimatedVector2) Controls(ui *lib.UIStruct, rect rl.Rectangle, comp components.Components) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		row, contrains := NewControlsLayout(rect.X, rect.Y, rect.Width)
		row.Add(Label("Position"))
		row.Add(v2.Inputs(ui, comp))
		return row.Draw, contrains
	}
}

func (v2 *AnimatedVector2) Inputs(ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		row := InputsLayout(2, rect)
		row.Add(v2.X.Input(ui, comp))
		row.Add(v2.Y.Input(ui, comp))
		row.Draw()
	}
}

func (v2 *AnimatedVector2) canDrawTimelineX() bool {
	return len(v2.X.SortedKeyframes) > 0
}
func (v2 *AnimatedVector2) canDrawTimelineY() bool {
	return len(v2.Y.SortedKeyframes) > 0
}
func (v2 *AnimatedVector2) CanDrawTimeline() bool {
	return v2.canDrawTimelineX() || v2.canDrawTimelineX()
}

func (v2 *AnimatedVector2) Timeline(ui *lib.UIStruct, comp components.Components) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := TimelinePanelRowLayout(rect)
		row.Add(TimelinePanelLabel("Position"))
		row.Add(v2.TimelineInputs(ui, comp))
		row.Draw()
		return row.Draw, row.CurrentRect.Width, row.CurrentRect.Height
	}
}
func (v2 *AnimatedVector2) TimelineInputs(ui *lib.UIStruct, comp components.Components) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := TimelinePanelInputsLayout(rect)
		if v2.canDrawTimelineX() {
			layout.Add(v2.X.NewInput(ui, comp))
		}
		if v2.canDrawTimelineY() {
			layout.Add(v2.Y.NewInput(ui, comp))
		}

		return layout.Draw, layout.CurrentRect.Width, layout.CurrentRect.Height
	}
}
