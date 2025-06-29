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

func (a *AnimatedColor) Controls(ui *lib.UIStruct, rect rl.Rectangle, comp components.Components) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		row, contrains := NewControlsLayout(avaliablePosition.X, avaliablePosition.Y, rect.Width)
		row.Add(Label("Color"))
		row.Add(a.ColorControlsInputs(ui, comp))
		return row.Draw, contrains
	}
}

func (a *AnimatedColor) ColorControlsInputs(ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		row := InputsLayout(4, rect)
		row.Add(a.red.Input(ui, comp))
		row.Add(a.green.Input(ui, comp))
		row.Add(a.blue.Input(ui, comp))
		row.Add(a.alpha.Input(ui, comp))
		row.Draw()
	}
}
