package components

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SidebarProperyLabel(text string) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		fontSize := 14
		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y+4, int32(fontSize), rl.White)
		}, 0, float32(fontSize)
	}
}

func NewSidebarProperyLabel(rect rl.Rectangle) *app.Layout {
	return app.
		NewLayout().
		PositionRect(rect).
		// TODO
		// BREAKS WHEN GAP IS AFTER WIDTH
		Gap(32).
		Width(rect.Width,
			app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 60},
			app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1}).
		Row()
}

func SidebrInputsLayout(amount int, rect rl.Rectangle) *app.Layout {
	childrenSize := make([]app.ChildSize, 0, amount)
	childSize := 1 / float32(amount)
	for i := 0; i < amount; i++ {
		childrenSize = append(childrenSize, app.ChildSize{
			SizeType: app.SIZE_WEIGHT,
			Value:    childSize,
		})
	}

	return app.
		NewLayout().
		PositionRect(rect).
		Row().
		Gap(8).
		Width(rect.Width, childrenSize...)
}
