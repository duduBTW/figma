package components

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type animatedVector2Component struct {
	prop   *app.AnimatedVector2
	layer  app.Layer
	prefix string

	x *animatedPropComponent
	y *animatedPropComponent
}

func NewAnimatedVector2(animatedProp *app.AnimatedVector2, layer app.Layer, prefix string) *animatedVector2Component {
	return &animatedVector2Component{
		prop:   animatedProp,
		layer:  layer,
		prefix: prefix,
		x:      NewAnimatedProp(&animatedProp.X, layer, prefix),
		y:      NewAnimatedProp(&animatedProp.Y, layer, prefix),
	}
}

func (component *animatedVector2Component) Controls() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := NewSidebarProperyLabel(rect).
			Add(SidebarProperyLabel("Position")).
			Add(component.colorControlsInputs())
		return row.Draw, 0, row.Size.Height
	}
}

func (component *animatedVector2Component) colorControlsInputs() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		l := component.layer
		row := SidebrInputsLayout(2, rect).
			Add(NewAnimatedProp(&component.prop.X, l, "").Input()).
			Add(NewAnimatedProp(&component.prop.Y, l, "").Input())
		return row.Draw, 0, row.Size.Height
	}
}

func (component *animatedVector2Component) Timeline() []app.Component {
	components := []app.Component{}
	if !component.x.CanDrawTimeline() && !component.y.CanDrawTimeline() {
		return components
	}

	if component.x.CanDrawTimeline() {
		components = append(components, TimelineRow("Position X", component.x.Input(), component.prop.X))
	}
	if component.y.CanDrawTimeline() {
		components = append(components, TimelineRow("Position Y", component.y.Input(), component.prop.Y))
	}
	return components
}
