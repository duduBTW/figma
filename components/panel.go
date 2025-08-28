package components

import (
	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SidebarProperyLabel(text string) app.Component {
	return Typography(text, ds.FONT_SIZE_LG, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT_ACCENT)
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

func NewPanelLayout(rect rl.Rectangle) *app.Layout {
	return app.
		NewLayout().
		PositionRect(rect).
		Width(rect.Width).
		Column().
		Gap(8)
}
