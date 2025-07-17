package layout

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type timelineLayout struct{}

func (timelineLayout) NewTilelineRowLayout(positionRect rl.Rectangle) *app.Layout {
	return app.
		NewLayout().
		PositionRect(positionRect).
		Row().
		Gap(12)
}

func (timelineLayout) Root(rect rl.Rectangle) *app.Layout {
	return app.
		NewLayout().
		PositionRect(rect).
		Column().
		Gap(12).
		Width(rect.Width)
}

func (timelineLayout) Row(rect rl.Rectangle) *app.Layout {
	return app.
		NewLayout().
		Row().
		PositionRect(rect).
		Gap(12).
		Width(rect.Width,
			app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 280},
			app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1})
}
func (timelineLayout) Panel(rect rl.Rectangle) *app.Layout {
	return app.
		NewLayout().
		PositionRect(rect).
		Padding(app.NewPadding().Start(32)).
		Row().
		Gap(16).
		Width(rect.Width,
			app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1},
			app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 100})
}

var Timeline = timelineLayout{}
