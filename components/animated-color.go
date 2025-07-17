package components

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type animatedColorComponent struct {
	prop   *app.AnimatedColor
	layer  app.Layer
	prefix string
}

func NewAnimatedColor(animatedProp *app.AnimatedColor, layer app.Layer, prefix string) *animatedColorComponent {
	return &animatedColorComponent{
		prop:   animatedProp,
		layer:  layer,
		prefix: prefix,
	}
}

func (component *animatedColorComponent) Controls() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := NewSidebarProperyLabel(rect).
			Add(SidebarProperyLabel("Color")).
			Add(component.colorControlsInputs())
		return row.Draw, 0, row.Size.Height
	}
}

func (c *animatedColorComponent) colorControlsInputs() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		l := c.layer
		row := SidebrInputsLayout(4, rect).
			Add(NewAnimatedProp(&c.prop.Red, l, "").Input()).
			Add(NewAnimatedProp(&c.prop.Green, l, "").Input()).
			Add(NewAnimatedProp(&c.prop.Blue, l, "").Input()).
			Add(NewAnimatedProp(&c.prop.Alpha, l, "").Input())
		return row.Draw, 0, row.Size.Height
	}
}
