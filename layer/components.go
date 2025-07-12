package layer

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewControlsLayout(rect rl.Rectangle) *lib.Layout {
	return lib.
		NewLayout().
		PositionRect(rect).
		// TODO
		// BREAKS WHEN GAP IS AFTER WIDTH
		Gap(32).
		Width(rect.Width,
			lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: 60},
			lib.ChildSize{SizeType: lib.SIZE_WEIGHT, Value: 1}).
		Row()
}

func NewPanelLayout(rect rl.Rectangle) *lib.Layout {
	return lib.
		NewLayout().
		PositionRect(rect).
		Width(rect.Width).
		Column().
		Gap(8)
}

func InputsLayout(amount int, rect rl.Rectangle) *lib.Layout {
	childrenSize := make([]lib.ChildSize, 0, amount)
	childSize := 1 / float32(amount)
	for i := 0; i < amount; i++ {
		childrenSize = append(childrenSize, lib.ChildSize{
			SizeType: lib.SIZE_WEIGHT,
			Value:    childSize,
		})
	}

	return lib.
		NewLayout().
		PositionRect(rect).
		Row().
		Gap(8).
		Width(rect.Width, childrenSize...)
}

func Label(text string) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		fontSize := 14
		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y+4, int32(fontSize), rl.White)
		}, 0, float32(fontSize)
	}
}

// ----------
// Timeline
// ----------

func TimelinePanelTitle(text string, layer Layer, ui *lib.UIStruct) lib.Component {
	var fontSize int32 = 16
	var height = float32(fontSize)
	return func(rect rl.Rectangle) (func(), float32, float32) {
		width := float32(rl.MeasureText(text, fontSize))
		interractable := components.NewInteractable(layer.GetName()+"panel-item", ui)
		if interractable.Event(rl.GetMousePosition(), rl.NewRectangle(rect.X, rect.Y, width, height)) {
			//
		}

		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, fontSize, rl.White)
		}, width, height
	}
}

func TimelinePanelInputsLayout(containerRect rl.Rectangle) *lib.Layout {
	return lib.
		NewLayout().
		PositionRect(containerRect).
		Column().
		Gap(12).
		Width(containerRect.Width)
}
