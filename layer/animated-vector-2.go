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
		row.Add(Inputs(v2, ui, comp))
		return row.Draw, contrains
	}
}

func Inputs(v2 *AnimatedVector2, ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		row := InputsLayout(2, rect)
		row.Add(v2.X.Input(ui, comp))
		row.Add(v2.Y.Input(ui, comp))
		row.Draw()
	}
}
