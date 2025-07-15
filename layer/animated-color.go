package layer

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedColor struct {
	red   AnimatedProp
	blue  AnimatedProp
	green AnimatedProp
	alpha AnimatedProp
}

const (
	red_KEY   = "red"
	blue_KEY  = "blue"
	green_KEY = "green"
	alpha_KEY = "alpha"
)

func NewAnimatedColor(red, blue, green, alpha float32) AnimatedColor {
	return AnimatedColor{
		red:   NewAnimatedProp(red, red_KEY),
		blue:  NewAnimatedProp(blue, blue_KEY),
		green: NewAnimatedProp(green, green_KEY),
		alpha: NewAnimatedProp(alpha, alpha_KEY),
	}
}

func (a *AnimatedColor) Get(selectedFrame int) rl.Color {
	red := uint8(a.red.KeyFramePosition(selectedFrame))
	green := uint8(a.green.KeyFramePosition(selectedFrame))
	blue := uint8(a.blue.KeyFramePosition(selectedFrame))
	alpha := uint8(a.alpha.KeyFramePosition(selectedFrame))
	return rl.NewColor(red, green, blue, alpha)
}

func (a *AnimatedColor) Controls(ui *lib.UIStruct, comp components.Components, layer Layer) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := NewControlsLayout(rect).
			Add(Label("Color")).
			Add(a.ColorControlsInputs(ui, comp, layer))
		return row.Draw, 0, row.Size.Height
	}
}

func (a *AnimatedColor) ColorControlsInputs(ui *lib.UIStruct, comp components.Components, layer Layer) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := InputsLayout(4, rect).
			Add(a.red.Input(ui, comp, layer, "")).
			Add(a.green.Input(ui, comp, layer, "")).
			Add(a.blue.Input(ui, comp, layer, "")).
			Add(a.alpha.Input(ui, comp, layer, ""))
		return row.Draw, 0, 0
	}
}
